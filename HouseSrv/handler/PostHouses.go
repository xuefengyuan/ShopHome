package handler

import (
    "context"

    example "ShopHome/HouseSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "github.com/garyburd/redigo/redis"
    "encoding/json"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
    "reflect"
)

// 发布房源
func (e *Example) PostHouses(ctx context.Context, req *example.PostHousesRequest, rsp *example.PostHousesResponse) error {
    beego.Info("PostHouses 发布房源信息 /api/v1.0/houses ")

    //创建返回空间
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

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

    var requestMap = make(map[string]interface{})
    json.Unmarshal(req.Max, &requestMap)
    for key, value := range requestMap {
        beego.Info(key, value)
    }


    house := models.House{}
    /* 插入房源信息 */
    house.Title = requestMap["title"].(string)
    // 价格
    price, _ := strconv.Atoi(requestMap["price"].(string))
    house.Price = price * 100

    house.Address = requestMap["address"].(string)
    house.Room_count, _ = strconv.Atoi(requestMap["room_count"].(string))

    house.Acreage, _ = strconv.Atoi(requestMap["acreage"].(string))
    house.Unit = requestMap["unit"].(string)

    house.Capacity, _ = strconv.Atoi(requestMap["capacity"].(string))
    house.Beds = requestMap["beds"].(string)
    // 押金
    deposit, _ := strconv.Atoi(requestMap["deposit"].(string))
    house.Deposit = deposit * 100

    house.Min_days, _ = strconv.Atoi(requestMap["min_days"].(string))
    house.Max_days, _ = strconv.Atoi(requestMap["max_days"].(string))

    // 设施
    // "facility":["1","2","3","7","12","14","16","17","18","21","22"]
    facility := []*models.Facility{}

    for _, fId := range requestMap["facility"].([]interface{}) {
        // 将设施编号转换成对应的类型
        f_Id, _ := strconv.Atoi(fId.(string))
        // 创建临时变量，使用设施编号创建的设施表对象指针
        temp := &models.Facility{Id: f_Id}
        facility = append(facility, temp)
    }

    // 添加地区
    area_id, _ := strconv.Atoi(requestMap["area_id"].(string))
    area := models.Area{Id: area_id}
    house.Area = &area
    // 添加user信息
    user := models.User{Id: userId}
    house.User = &user

    o := orm.NewOrm()
    houserId, err := o.Insert(&house)
    if err != nil {
        beego.Info("房源信息插入失败,err = ",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    beego.Info(houserId,reflect.TypeOf(houserId),house.Id)
    /* 插入到房源与设施信息的多对多表中 */
    m2m := o.QueryM2M(&house, "Facilities")

    num, err := m2m.Add(facility)
    if err != nil {
        beego.Info("房屋设施插入失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    if num == 0 {
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    rsp.HouseId = int64(house.Id)

    return nil
}
