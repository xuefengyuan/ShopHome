package main

import (
    "github.com/micro/go-log"
    "net/http"
    _ "ShopHome/HomeWeb/modules"
    "github.com/micro/go-web"
    "github.com/julienschmidt/httprouter"
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
    // register html handler

    // register call handler
    service.Handle("/", rou)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
