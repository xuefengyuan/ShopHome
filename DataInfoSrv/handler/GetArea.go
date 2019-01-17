package handler

import (
    "context"

    example "ShopHome/DataInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "encoding/json"
    "github.com/astaxie/beego/cache"

    _"github.com/astaxie/beego/cache/redis"
    _"github.com/garyburd/redigo/redis"
    _"github.com/gomodule/redigo/redis"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
)

type Example struct{}

func RedisInit() (cache.Cache, error) {
    // 连接Redis,配置缓存参数
    redisConf := map[string]string{
        "key": utils.G_server_name,
        //127.0.0.1:6379
        "conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
        "dbNum": utils.G_redis_dbnum,
    }
    beego.Info(redisConf)

    // 将map进行转化成为json
    redisConfJs, _ := json.Marshal(redisConf)

    // 创建redis句柄
    bm, err := cache.NewCache("redis", string(redisConfJs))

    return bm, err
}


// 获取地区信息服务
func (e *Example) GetArea(ctx context.Context, req *example.DataInfoRequest, rsp *example.DataInfoResponse) error {
    beego.Info("请求地区信息 GetArea api/v1.0/areas")
    // 初始化 错误码
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    // 创建Redis句柄
    bm, err := RedisInit()
    if err != nil {
        beego.Info("redis连接失败",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
    }

    // 获取数据 在这里我们需要定制1个key 就算用来作area查询的  area_info
    areaValue := bm.Get(utils.AREA_INFO_KEY)

    if areaValue != nil {
        /*如果有数据就发送给前端*/
        beego.Info("获取到地域信息缓存")
        areaMap := []map[string]interface{}{}
        // 将获取到的数据进行json的解码操作
        json.Unmarshal(areaValue.([]byte),&areaMap)

        beego.Info("得到从缓存中提取的area数据 ",areaMap)

        for _,value := range areaMap{
            tmp := example.DataInfoResponse_Areas{Aid:int32(value["aid"].(float64)),Aname:value["aname"].(string)}
            rsp.Data = append(rsp.Data,&tmp)
        }
        // 将数据发送给了web后面就不需要执行了
        return nil
    }

    /*2没有数据就从mysql中查找数据*/
    // beego 操作数据库的orm方法
    // 创建orm句柄
    o := orm.NewOrm()
    // 查询什么
    qs := o.QueryTable("area")
    // 用什么接收
    var areas []models.Area
    num, err := qs.All(&areas)
    if err != nil {
        beego.Info("数据库查询失败", err)
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    if num == 0 {
        beego.Info("数据库没有数据", num)
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }


    /*3将查找到的数据存到缓存中*/
    // 需要将获取到的数据转化为json
    areaJson,_ := json.Marshal(areas)
    err = bm.Put(utils.AREA_INFO_KEY, areaJson, time.Second*3600)

    if err != nil {
        beego.Info("数据缓存失败",err)
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
    }

    /*4将查找到的数据发送给前端*/
    for _, value := range areas {
        beego.Info("key = ", value)
        tmp := example.DataInfoResponse_Areas{Aid: int32(value.Id), Aname: value.Name}
        rsp.Data = append(rsp.Data, &tmp)
    }

    return nil
}


