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

    // 获取验证码图片
    rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCode)
    // 获取短信验证码
    rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmsCode)

    // 用户注册
    rou.POST("/api/v1.0/users", handler.PostUserRegist)
    // 用户登陆
    rou.POST("/api/v1.0/sessions", handler.PostUserLogin)

    // 获取用户信息
    rou.GET("/api/v1.0/user", handler.GetUserInfo)
    // 用户退出登陆
    rou.DELETE("/api/v1.0/session", handler.DeleteSession)


    // 上传头像
    rou.POST("/api/v1.0/user/avatar", handler.PostUserAvatar)
    // 请求更新用户名 PUT
    rou.PUT("/api/v1.0/user/name", handler.PutUserInfo)

    // 实名认证检查 GET，调用获取用户信息方法
    rou.GET("/api/v1.0/user/auth", handler.GetUserInfo)
    // 实名认证 post
    rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)



    // register call handler
    service.Handle("/", rou)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
