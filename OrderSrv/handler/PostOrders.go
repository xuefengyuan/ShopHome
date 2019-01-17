package handler

import (
    "context"

    example "ShopHome/OrderSrv/proto/example"
    "ShopHome/HomeWeb/utils"
    "github.com/astaxie/beego"
    "encoding/json"
    "github.com/astaxie/beego/cache"
    _ "github.com/astaxie/beego/cache/redis"
    _ "github.com/garyburd/redigo/redis"
    _ "github.com/gomodule/redigo/redis"
    "ShopHome/HomeWeb/modules"
    "strconv"
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/garyburd/redigo/redis"
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

// 发布订单
func (e *Example) PostOrders(ctx context.Context, req *example.PostOrdersRequest, rsp *example.PostOrdersResponse) error {
    beego.Info("PostOrders  发布订单 /api/v1.0/orders")
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    // 1得到用户请求的json数据并效验合法性
    // 获取用户请求Response数据的name
    var requestMap = make(map[string]interface{})
    err := json.Unmarshal(req.Body, &requestMap)

    if err != nil {
        beego.Info("body 转码失败", err)
        rsp.Errno = utils.RECODE_REQERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil

    }
    beego.Info("requestMap = ", requestMap)
    // 效验合法性
    if requestMap["house_id:"] == "" || requestMap["start_date"] == "" || requestMap["end_date"] == "" {
        beego.Info("数据校验失败", err)
        rsp.Errno = utils.RECODE_REQERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }


    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }


    // 创建redis句柄
    bm, err := RedisInit()

    /* 从从sessionid获取当前的userid */
    sessionId := req.SessionId

    sessionUserId := bm.Get(sessionId + "user_id")
    userIdStr, _ := redis.String(sessionUserId, nil)
    userId, _ := strconv.Atoi(userIdStr)

    beego.Info("userid : ", userId)

    // 确定end_date在start_data之后
    startDataTime, _ := time.Parse("2006-01-02 15:04:05", requestMap["start_date"].(string)+" 00:00:00")
    endDataTime, _ := time.Parse("2006-01-02 15:04:05", requestMap["end_date"].(string)+" 00:00:00")
    //得到一共入住的天数

    beego.Info(startDataTime, endDataTime)
    days := endDataTime.Sub(startDataTime).Hours()/24 + 1
    beego.Info("days = ", days)

    // 根据order_id得到关联的房源信息
    houseId, _ := strconv.Atoi(requestMap["house_id"].(string))
    // 房屋对象
    house := models.House{Id: houseId}

    o := orm.NewOrm()
    if err := o.Read(&house); err != nil {
        beego.Info("查询房屋失败", err)
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    // 关联查询用户
    o.LoadRelated(&house, "User")
    // 确保当前的uers_id不是房源信息所关联的user_id
    if userId == house.User.Id{
        beego.Info("房东自己租房啦")
        rsp.Errno  =  utils.RECODE_ROLEERR
        rsp.ErrMsg  = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 确保用户选择的房屋未被预定,日期没有冲突
    if endDataTime.Before(startDataTime){
        rsp.Errno  =  utils.RECODE_ROLEERR
        rsp.ErrMsg  = "结束时间在开始时间之前"
        return nil
    }

    // TODO 添加征信步骤,暂未处理

    // 封装order订单
    amount := days* float64(house.Price)

    order := models.OrderHouse{}
    order.House = &house
    user := models.User{Id: userId}
    order.User = &user

    order.Begin_date = startDataTime
    order.End_date = endDataTime

    order.Days = int(days)
    order.House_price = house.Price

    order.Amount = int(amount)
    // 订单状态，初始为待确认
    order.Status = models.ORDER_STATUS_WAIT_ACCEPT
    // 征信
    order.Credit = true
    beego.Info("order = ",order)

    // 将订单信息入库表中
    if _, err := o.Insert(&order); err != nil{
        rsp.Errno  =  utils.RECODE_DBERR
        rsp.ErrMsg  = utils.RecodeText(rsp.Errno)
        return nil
    }
    // 更新缓存
    bm.Put(sessionId+"user_id", userId, time.Second*3600)

    // 返回订单id
    rsp.OrderId = int64(order.Id)

    return nil
}
