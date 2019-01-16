package main

import (
    "github.com/micro/go-log"
    "net/http"
    _ "ShopHome/HomeWeb/modules"
    "github.com/micro/go-web"
    "github.com/julienschmidt/httprouter"
    "ShopHome/HomeWeb/handler"
)

func main() {
    // create new web service
    service := web.NewService(
        web.Name("go.micro.web.HomeWeb"),
        web.Version("latest"),
        web.Address(":8991"),
    )

    // initialise service
    if err := service.Init(); err != nil {
        log.Fatal(err)
    }

    rou := httprouter.New()
    rou.NotFound = http.FileServer(http.Dir("html"))

    // 获取地区请求
    rou.GET("/api/v1.0/areas", handler.GetArea)
    // 获取session
    rou.GET("/api/v1.0/session", handler.GetSession)

    // 获取首页轮播图
    rou.GET("/api/v1.0/house/index", handler.GetIndex)


    // register call handler
    service.Handle("/", rou)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
