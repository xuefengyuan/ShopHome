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
    "fmt"
)


// 用户评价订单信息
func (e *Example) PutComment(ctx context.Context, req *example.UserCommentRequest, rsp *example.UserCommentResponse) error {
    beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")

    // 创建返回空间
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

    sessionUserId := bm.Get(sessionId + "user_id")
    userIdStr, _ := redis.String(sessionUserId, nil)

    userId, _ := strconv.Atoi(userIdStr)
    beego.Info("userid : ", userId)

    orderId, _ := strconv.Atoi(req.OrderId)
    comment := req.Comment
    // 检验评价信息是否合法 确保不为空
    if comment == "" {
        rsp.Errno = utils.RECODE_PARAMERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    order := models.OrderHouse{}
    o := orm.NewOrm()
    beego.Info("order id = ",orderId)

    err  = o.QueryTable("order_house").Filter("id", orderId).Filter("status", models.ORDER_STATUS_WAIT_COMMENT).One(&order)

    if err != nil {
        beego.Info("查询房屋订单信息失败 ",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 关联查询order订单所关联的user信息
    if _, err := o.LoadRelated(&order, "User"); err != nil {

        beego.Info("关联查询订单用户失败 ",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    // 确保订单所关联的用户和该用户是同一个人
    if userId != order.User.Id {
        beego.Info("用户和订单不是同一人")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 关联查询order订单所关联的House信息

    if _, err := o.LoadRelated(&order, "house"); err != nil {
        beego.Info("关联查询房屋信息失败 ",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
    }

    house := order.House

    // 将房源信息的评论字段追加评论信息
    // 更新order的status为COMPLETE
    order.Status = models.ORDER_STATUS_COMPLETE
    order.Comment = comment

    // 将房屋订单成交量+1
    house.Order_count++

    // 将order和house更新数据库
    if _, err := o.Update(&order, "status", "comment"); err != nil {
        beego.Info("订单信息更新失败 ",err)
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    if _,err := o.Update(&house,"order_count");err != nil{
        beego.Info("房屋信息更新失败 ",err)
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    // 将house_info_[house_id]的缓存key删除 （因为已经修改订单数量）
    houseInfoKey := fmt.Sprint("house_info_%s",house.Id)
    bm.Delete(houseInfoKey)

    return nil
}
