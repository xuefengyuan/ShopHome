package handler

import (
    "context"
    "encoding/json"
    "net/http"
    ORDER "ShopHome/OrderSrv/proto/example"
    "github.com/julienschmidt/httprouter"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "io/ioutil"
    "github.com/micro/go-grpc"
)

// 发布订单
func PostOrders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("PostOrders  发布订单 /api/v1.0/orders")

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

    //将post传过来的数据转化以下
    body, _ := ioutil.ReadAll(r.Body)

    beego.Info("body = ", body)

    server := grpc.NewService()
    server.Init()

    // call the backend service
    exampleClient := ORDER.NewExampleService("go.micro.srv.OrderSrv", server.Client())
    rsp, err := exampleClient.PostOrders(context.TODO(), &ORDER.PostOrdersRequest{
        SessionId: cookie.Value,
        Body:      body,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    /* 得到插入房源信息表的 id */
    houserIdMap := make(map[string]interface{})
    houserIdMap["order_id"] = int(rsp.OrderId)

    // 返回数据给前端
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   houserIdMap,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

// 查看房东/租客订单的
func GetUserOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info(" 查看订单信息（房东/租客） GetUserOrder  /api/v1.0/user/orders ")
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

    server := grpc.NewService()
    server.Init()

    role := r.URL.Query()["role"][0]

    // call the backend service
    exampleClient := ORDER.NewExampleService("go.micro.srv.OrderSrv", server.Client())
    rsp, err := exampleClient.GetUserOrder(context.TODO(), &ORDER.UserOrderRequest{
        SessionId: cookie.Value,
        Role:      role,
    })
    if err != nil {
        beego.Info("err = ",err)
        http.Error(w, err.Error(), 500)
        return
    }

    // 接收数据
    orderList := []interface{}{}
    json.Unmarshal(rsp.Orders, &orderList)

    data := map[string]interface{}{}
    data["orders"] = orderList

    // 返回数据给前端
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

// 房东同意/拒绝订单的
func PutOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    beego.Info(" 房东同意/拒绝订单 PutOrders  /api/v1.0/orders/:id/status ")

    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
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

    server := grpc.NewService()
    server.Init()

    // call the backend service
    exampleClient := ORDER.NewExampleService("go.micro.srv.OrderSrv", server.Client())
    rsp, err := exampleClient.PutOrders(context.TODO(), &ORDER.PutOrdersRequest{
        SessionId: cookie.Value,
        Action:    request["action"].(string),
        OrderId:   ps.ByName("id"),
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 返回数据给前端
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

// 用户评价订单
func PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")
    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
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

    server := grpc.NewService()
    server.Init()
    exampleClient := ORDER.NewExampleService("go.micro.srv.OrderSrv", server.Client())
    rsp, err := exampleClient.PutComment(context.TODO(), &ORDER.UserCommentRequest{
        SessionId: cookie.Value,
        Comment:   request["comment"].(string),
        OrderId:   ps.ByName("id"),
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 返回数据给前端
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
