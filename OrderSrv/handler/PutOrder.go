package handler

import (
    "context"

    example "ShopHome/OrderSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "ShopHome/HomeWeb/modules"
    "strconv"
    "github.com/astaxie/beego/orm"
    "github.com/garyburd/redigo/redis"
)


// 房东同意/拒绝订单的
func (e *Example) PutOrders(ctx context.Context, req *example.PutOrdersRequest, rsp *example.PutOrdersResponse) error {
    beego.Info(" 房东同意/拒绝订单 PutOrders  /api/v1.0/orders/:id/status ")
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

    /* 从从sessionid获取当前的userid */
    sessionId := req.SessionId


    sessionUserId := bm.Get(utils.GetUserIdKey(sessionId))
    userIdStr, _ := redis.String(sessionUserId, nil)

    userId, _ := strconv.Atoi(userIdStr)
    beego.Info("userid : ", userId)

    orderId, _ := strconv.Atoi(req.OrderId)
    // 解析客户端请求的json数据得到action参数
    beego.Info(req.Action)
    action := req.Action

    beego.Info("action = ", action)

    order := models.OrderHouse{}
    o := orm.NewOrm()
    err  = o.QueryTable("order_house").Filter("id", orderId).Filter("status", models.ORDER_STATUS_WAIT_ACCEPT).One(&order)

    if err != nil {
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    beego.Info("order id ",order.Id)

    // 关联查询房屋信息
    if _, err := o.LoadRelated(&order, "House"); err != nil {
        rsp.Errno  =  utils.RECODE_DATAERR
        rsp.ErrMsg  = utils.RecodeText(rsp.Errno)
        return nil
    }

    beego.Info("order house id ",order.House.Id)

    house := order.House
    // 校验该订单的user_id是否是当前用户的user_id
    if house.User.Id != userId {
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = "订单用户不匹配,操作无效"
        return nil
    }

    if action == "accept" {
        // 如果是接受订单,将订单状态变成待评价状态
        order.Status = models.ORDER_STATUS_WAIT_COMMENT
        beego.Info("同意订单")

    } else if action == "reject" {
        // 如果是拒绝接单, 尝试获得拒绝原因,并把拒单原因保存起来
        order.Status = models.ORDER_STATUS_REJECTED

        // 更换订单状态为status为reject
        reason := req.Action
        order.Comment = reason
        beego.Info("拒绝订单")
    }

    // 更新该数据到数据库中
    if _, err := o.Update(&order); err != nil {
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
    }

    return nil
}
