package handler

import (
    "context"

    example "ShopHome/DataInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "encoding/json"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
)

// 获取首页轮播图
func (e *Example) GetIndex(ctx context.Context, req *example.IndexRequest, rsp *example.IndexResponse) error {
    beego.Info("获取首页轮播图信息 GetSession /api/v1.0/houses/index")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    // 创建redis句柄
    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    housePageValue := bm.Get(utils.HOME_INDEX_DATA_KEY)
    if housePageValue != nil {
        // 直接将二进制发送给客户端
        rsp.Max = housePageValue.([]byte)
        return nil
    }

    data := []interface{}{}

    houses := []models.House{}
    // 2 如果缓存没有,需要从数据库中查询到房屋列表
    o := orm.NewOrm()

    if _, err = o.QueryTable("house").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses); err == nil {
        for _, house := range houses {
            o.LoadRelated(&house, "User")
            o.LoadRelated(&house, "Area")
            o.LoadRelated(&house, "Facilities")
            o.LoadRelated(&house, "Images")
            // TODO 订单暂时没有， 先不查询
            //o.LoadRelated(&house,"Orders")
            // 房屋信息转换
            data = append(data, house.To_house_info())
        }
    }
    beego.Info("data : ", data)

    housePageValue, _ = json.Marshal(data)
    // 将data存入缓存数据
    bm.Put(utils.HOME_INDEX_DATA_KEY, housePageValue, time.Second*3600)

    rsp.Max = housePageValue.([]byte)

    return nil
}
