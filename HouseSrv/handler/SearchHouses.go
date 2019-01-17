package handler

import (
    "context"

    example "ShopHome/HouseSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
    "encoding/json"
)

// 搜索房源
func (e *Example) GetSearchHouses(ctx context.Context, req *example.SearchHousesRequest, rsp *example.SearchHousesResponse) error {
    beego.Info("搜索房源详细信息 GetHouses  /api/v1.0/houses ")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    aid, _ := strconv.Atoi(req.Aid)
    sd := req.Sd
    ed := req.Ed
    sk := req.Sk  // 第三栏的信息
    page,_ := strconv.Atoi(req.P)
    // 可以根据上面参数做其它复杂条件过滤
    beego.Info(aid, sd, ed, sk, page)


    houses := []models.House{}
    o := orm.NewOrm()
    // 设置查找的表
    qs := o.QueryTable("house")
    // 根据查询条件来查找内容
    // 查找传入地区的所有房屋
    num, err := qs.Filter("area_id", aid).All(&houses)

    if err != nil {
        rsp.Errno = utils.RECODE_PARAMERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    if num == 0 {
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    beego.Info("num : ",num)
    // 计算一下 所有房屋 / 一页现实的数量
    // 计算房源总页数，查询的房屋数量/每页显示数量+1
    totalPage := int(num)/models.HOUSE_LIST_PAGE_CAPACITY + 1
    housePage := 1

    houseList := []interface{}{}
    for _,house := range houses {
        // 关联查找
        o.LoadRelated(&house,"Area")
        o.LoadRelated(&house,"User")
        o.LoadRelated(&house,"Images")
        o.LoadRelated(&house,"Facilities")
        houseList = append(houseList, house.To_house_info())
    }

    // 返回数据
    rsp.TotalPage = int64(totalPage)
    rsp.CurrentPage = int64(housePage)
    // 数据编码返回web服务
    rsp.Houses,_  = json.Marshal(houseList)
    beego.Info("totalPage = ",totalPage)

    return nil
}
