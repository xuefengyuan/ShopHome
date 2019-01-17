package handler

import (
    "context"

    example "ShopHome/HouseSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "github.com/garyburd/redigo/redis"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
    "encoding/json"
    "time"
    "fmt"
)

// 获取房源详情
func (e *Example) GetHouseInfo(ctx context.Context, req *example.HouseInfoRequest, rsp *example.HouseInfoResponse) error {
    beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")
    //创建返回空间
    rsp.Errno  =  utils.RECODE_OK
    rsp.ErrMsg  = utils.RecodeText(rsp.Errno)

    // 创建redis句柄
    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 从从sessionid获取当前的userid */
    sessionId := req.SessionId
    sessionUserId := bm.Get(utils.GetUserIdKey(sessionId))
    userIdStr, _ := redis.String(sessionUserId, nil)
    userId, _ := strconv.Atoi(userIdStr)

    beego.Info("userid : ", userId)

    /* 从请求中的url获取房源id */
    houseId, _ := strconv.Atoi(req.HouseId)
    /* 从缓存数据库中获取到当前房屋的数据 */

    houseInfoKey := fmt.Sprint("house_info_%s",houseId)
    houseInfoValue := bm.Get(houseInfoKey)

    // 有缓存数据就直接返回了
    if houseInfoValue != nil {
        rsp.UserId = int64(userId)
        rsp.HouseData = houseInfoValue.([]byte)
        return nil
    }

    house := models.House{Id: houseId}

    o := orm.NewOrm()

    err = o.Read(&house)
    if err != nil {
        beego.Info("查询房源失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 关联查询 area user images fac等表 */

    o.LoadRelated(&house,"Area")
    o.LoadRelated(&house,"User")
    o.LoadRelated(&house,"Images")
    o.LoadRelated(&house,"Facilities")
    // TODO 暂时没有订单
    // o.LoadRelated(&house,"Orders")
    /* 将查询到的结果存储到缓存当中 */
    houseMix, _:= json.Marshal(house)
    bm.Put(houseInfoKey,houseMix,time.Second*3600)

    rsp.UserId = int64(userId)
    rsp.HouseData = houseMix


    return nil
}
