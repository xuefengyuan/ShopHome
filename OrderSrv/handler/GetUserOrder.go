package handler

import (
    "context"

    example "ShopHome/OrderSrv/proto/example"
    "ShopHome/HomeWeb/utils"
    "github.com/astaxie/beego"
    "strconv"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
    "encoding/json"
    "github.com/garyburd/redigo/redis"
)

// 查看房东/租客订单的
func (e *Example) GetUserOrder(ctx context.Context, req *example.UserOrderRequest, rsp *example.UserOrderResponse) error {
    beego.Info(" 查看订单信息（房东/租客） GetUserOrder  /api/v1.0/user/orders ")
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

    sessionId := req.SessionId

    tmpid := bm.Get(utils.GetUserIdKey(sessionId))
    userId,_:=strconv.Atoi(string(tmpid.([]byte)))


    beego.Info("user id ", userId)

    userIdStr, _ := redis.String(userId, nil)
    beego.Info("user id str ", userIdStr)
    id, _ := strconv.Atoi(userIdStr)
    beego.Info("user id int ", id)

    // 得到用户角色
    beego.Info("role = ", req.Role)

    o := orm.NewOrm()

    orders := []models.OrderHouse{}
    // 存放订单的切片
    orderList := []interface{}{}

    // 标识是房东还是租客可看，role=custom：租客查看，role=landlord房东查看
    if "landlord" == req.Role {
        //角色为房东
        //现找到自己目前已经发布了哪些房子
        landLordHouses := []models.House{}
        o.QueryTable("house").Filter("user__id", userId).All(&landLordHouses)

        housesIds := []int{}
        for _, house := range landLordHouses {
            housesIds = append(housesIds, house.Id)
        }
        //在从订单中找到房屋id为自己房源的id
        //o.QueryTable("order_house").Filter("house__id__in", housesIds).All(&orders)
        o.QueryTable("order_house").Filter("house__id__in", housesIds).OrderBy("ctime").All(&orders)
    } else {
        //角色为租客
        //_,err:=o.QueryTable("order_house").Filter("user__id", userId).All(&orders)
        _,err:=o.QueryTable("order_house").Filter("user__id", userId).OrderBy("ctime").All(&orders)
        if err != nil {
            beego.Info(err)
        }

    }


    // 循环将数据放到切片中
    for _, order := range orders {
        o.LoadRelated(&order,"User")
        o.LoadRelated(&order,"House")
        orderList = append(orderList,order.To_order_info())
    }

    beego.Info("order list = ",orderList)

    rsp.Orders ,_ = json.Marshal(orderList)

    return nil
}
