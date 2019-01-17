package handler

import (
    "context"
    "encoding/json"
    "net/http"
    DATAINFO "ShopHome/DataInfoSrv/proto/example"
    "github.com/julienschmidt/httprouter"
    "github.com/astaxie/beego"
    "github.com/micro/go-grpc"
    "ShopHome/HomeWeb/modules"
)

/* 获取地区信息 */
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("请求地区信息 GetArea api/v1.0/areas")

    // 创建服务获取句柄
    server := grpc.NewService()
    // 初始化服务
    server.Init()
    // 调用服务返回句柄
    exampleClient := DATAINFO.NewExampleService("go.micro.srv.DataInfoSrv", server.Client())
    // 调用服务方法
    rsp, err := exampleClient.GetArea(context.TODO(), &DATAINFO.DataInfoRequest{})
    if err != nil {
        beego.Info("err ：",err)
        http.Error(w, err.Error(), 500)
        return
    }

    // 接收数据
    // 准备接收切片
    areaList := []models.Area{}

    // 循环接收数据
    for _,value := range rsp.Data{
        area := models.Area{Id: int(value.Aid), Name: value.Aname}
        areaList = append(areaList,area)
    }

    // // 返回给前端的map
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data": areaList,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 获取首页轮播图 */
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("获取首页轮播图信息 GetSession /api/v1.0/houses/index")

    server := grpc.NewService()
    server.Init()
    // 连接服务
    exampleClient := DATAINFO.NewExampleService("go.micro.srv.DataInfoSrv", server.Client())
    rsp, err := exampleClient.GetIndex(context.TODO(), &DATAINFO.IndexRequest{})
    if err != nil {
        beego.Info("err ", err)
        http.Error(w, err.Error(), 500)
        return
    }

    data := []interface{}{}
    json.Unmarshal(rsp.Max, &data)

    beego.Info("data", data)
    //创建返回数据map
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
