package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/micro/go-micro/client"
    example "github.com/micro/examples/template/srv/proto/example"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
    // decode the incoming request as json
    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // call the backend service
    exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
    rsp, err := exampleClient.Call(context.TODO(), &example.Request{
        Name: request["name"].(string),
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // we want to augment the response
    response := map[string]interface{}{
        "errno": rsp.Msg,
        "errmsg": time.Now().UnixNano(),
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
