package handler

import (
    "context"

    example "ShopHome/HouseSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "encoding/json"
    "github.com/astaxie/beego/cache"
    _ "github.com/astaxie/beego/cache/redis"
    _ "github.com/garyburd/redigo/redis"
    _ "github.com/gomodule/redigo/redis"
    "strconv"
    "github.com/garyburd/redigo/redis"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
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

// 获取用户的房源信息
func (e *Example) GetUserHouses(ctx context.Context, req *example.UserHousesRequest, rsp *example.UserHousesResponse) error {
    beego.Info("获取当前用户所发布的房源 GetUserHouses /api/v1.0/user/houses")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    sessionId := req.SessionId

    userId := bm.Get(utils.GetUserIdKey(sessionId))
    beego.Info("user id ", userId)

    userIdStr, _ := redis.String(userId, nil)
    beego.Info("user id str ", userIdStr)

    id, _ := strconv.Atoi(userIdStr)
    beego.Info("user id int ", id)

    /* 通过user_id 获取到当前的用户所发布的房源信息 */
    houseList := []models.House{}
    // 创建数据库句柄
    o := orm.NewOrm()
    qs := o.QueryTable("house")

    // 过滤条件，用户的id
    num, err := qs.Filter("user__id", id).All(&houseList)

    if err != nil {
        beego.Info("房源数据库查询失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    if num == 0 {
        beego.Info("房源数据库查询 ： ", num)
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 成功返回数据给前端 */
    house, _ := json.Marshal(houseList)
    rsp.Mix = house

    return nil
}
