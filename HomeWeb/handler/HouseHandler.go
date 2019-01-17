package handler

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    HOUSESRV "ShopHome/HouseSrv/proto/example"
    "github.com/micro/go-grpc"
    "context"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "ShopHome/HomeWeb/modules"
    "fmt"
    "io/ioutil"
)

/* 获取用户发布的房源 */
func GetUserHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("获取当前用户所发布的房源 GetUserHouses /api/v1.0/user/houses")

    cookie, err := r.Cookie("userlogin")

    beego.Info("session id : ")
    if err != nil || cookie == nil || cookie.Value == "" {
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }
        beego.Info("err  1 : ", err)
        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := HOUSESRV.NewExampleService("go.micro.srv.HouseSrv", server.Client())
    rsp, err := exampleClient.GetUserHouses(context.TODO(), &HOUSESRV.UserHousesRequest{
        SessionId: cookie.Value,
    })
    if err != nil {
        beego.Info("err : ", err)
        http.Error(w, err.Error(), 500)
        return
    }

    houseList := []models.House{}
    // json解析数据
    json.Unmarshal(rsp.Mix, &houseList)

    var houses []interface{}
    // 遍历房源信息
    for _, houseInfo := range houseList {
        fmt.Printf("house.user = %+v\n", houseInfo.Id)
        fmt.Printf("house.area = %+v\n", houseInfo.Area)
        houses = append(houses, houseInfo.To_house_info())
    }

    dataMap := make(map[string]interface{})
    dataMap["houses"] = houses
    // 返回给前端的map
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   dataMap,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 发布房源 */
func PostHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("PostHouses 发布房源信息 /api/v1.0/houses ")

    cookie, err := r.Cookie("userlogin")

    if err != nil || cookie == nil || cookie.Value == "" {
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }
        beego.Info("err  1 : ", err)
        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    server := grpc.NewService()
    server.Init()

    // 获取前端post请求发送的内容
    body, _ := ioutil.ReadAll(r.Body)

    // call the backend service
    exampleClient := HOUSESRV.NewExampleService("go.micro.srv.HouseSrv", server.Client())
    rsp, err := exampleClient.PostHouses(context.TODO(), &HOUSESRV.PostHousesRequest{
        SessionId: cookie.Value,
        Max:       body,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    /* 得到插入房源信息表的 id */
    houseIdMap := make(map[string]interface{})
    houseIdMap["house_id"] = int(rsp.HouseId)

    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   houseIdMap,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 上传房屋图片 */
func PostHousesImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    beego.Info("发送房屋图片 PostHousesImage  /api/v1.0/houses/:id/images")

    // 获取到前端发送的文件信息

    file, fileHeader, err := r.FormFile("house_image")
    if err != nil {

        response := map[string]interface{}{
            "errno":  utils.RECODE_IOERR,
            "errmsg": utils.RecodeText(utils.RECODE_IOERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    beego.Info("文件大小", fileHeader.Size)
    beego.Info("文件名", fileHeader.Filename)

    // 创建一个文件大小的切片
    fileBuf := make([]byte, fileHeader.Size)
    // 将file的数据读到filebuf
    _, err = file.Read(fileBuf)

    if err != nil {
        beego.Info("文件读取异常：", err)
        response := map[string]interface{}{
            "errno":  utils.RECODE_IOERR,
            "errmsg": utils.RecodeText(utils.RECODE_IOERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    // 获取cookie
    cookie, err := r.Cookie("userlogin")

    if err != nil || cookie == nil || cookie.Value == "" {
        beego.Info("cookie 读取异常：", err)
        response := map[string]interface{}{
            "errno":  utils.RECODE_DATAERR,
            "errmsg": utils.RecodeText(utils.RECODE_DATAERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    houseid := ps.ByName("id")

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := HOUSESRV.NewExampleService("go.micro.srv.HouseSrv", server.Client())
    rsp, err := exampleClient.PostHousesImage(context.TODO(), &HOUSESRV.HousesImageRequest{
        SessionId: cookie.Value,
        HouseId:   houseid,
        Image:     fileBuf,
        FileName:  fileHeader.Filename,
        FileSize:  fileHeader.Size,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    data := make(map[string]interface{})
    data["url"] = utils.AddDomain2Url(rsp.ImgUrl)

    beego.Info("data: ", data)
    // 返回数据Map
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   data,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 获取房源详细信息 */
func GetHouseInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

    // 获取cookie
    cookie, err := r.Cookie("userlogin")

    if err != nil || cookie == nil || cookie.Value == "" {
        beego.Info("cookie 读取异常：", err)
        response := map[string]interface{}{
            "errno":  utils.RECODE_DATAERR,
            "errmsg": utils.RecodeText(utils.RECODE_DATAERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    houseid := ps.ByName("id")

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := HOUSESRV.NewExampleService("go.micro.srv.HouseSrv", server.Client())
    rsp, err := exampleClient.GetHouseInfo(context.TODO(), &HOUSESRV.HouseInfoRequest{
        SessionId: cookie.Value,
        HouseId:   houseid,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    house := models.House{}
    json.Unmarshal(rsp.HouseData, &house)

    dataMap := make(map[string]interface{})
    dataMap["user_id"] = rsp.UserId
    dataMap["house"] = house.To_one_house_desc()

    // 返回数据给前端
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   dataMap,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 搜索房源 */
func GetSearchHouses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    beego.Info("搜索房源详细信息 GetHouses  /api/v1.0/houses ")

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := HOUSESRV.NewExampleService("go.micro.srv.HouseSrv", server.Client())

    // aid=5&sd=2017-11-12&ed=2017-11-30&sk=new&p=1
    aid := r.URL.Query()["aid"][0] // 地区编号
    sd := r.URL.Query()["sd"][0]   // 开始时间
    ed := r.URL.Query()["ed"][0]   // 结束时间
    sk := r.URL.Query()["sk"][0]   // 第三栏条件
    p := r.URL.Query()["p"][0]     // 页码
    beego.Info("adi = ", aid)

    rsp, err := exampleClient.GetSearchHouses(context.TODO(), &HOUSESRV.SearchHousesRequest{
        Aid: aid,
        Sd:  sd,
        Ed:  ed,
        Sk:  sk,
        P:   p,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 接收服务返回的数据
    houses := []interface{}{}
    // 解码数据
    json.Unmarshal(rsp.Houses, &houses)
    // 拼装返回前端的数据
    data := map[string]interface{}{}
    data["current_page"] = rsp.CurrentPage
    data["houses"] = houses
    data["total_page"] = rsp.TotalPage

    beego.Info("TotalPage : ", rsp.TotalPage)

    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   data,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
