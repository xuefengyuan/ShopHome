package handler

import (
    "context"
    "encoding/json"
    "net/http"
    USERINFOSRV "ShopHome/UserInfoSrv/proto/example"
    "github.com/julienschmidt/httprouter"
    "ShopHome/HomeWeb/utils"
    "github.com/astaxie/beego"
    "github.com/micro/go-grpc"
)

/* 用户注册 */
func PostUserRegist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("PostRet  注册 /api/v1.0/users")

    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    beego.Info(request)
    // 数据校验
    if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {

        response := map[string]interface{}{
            "errno":  utils.RECODE_NODATA,
            "errmsg": utils.RecodeText(utils.RECODE_DATAERR),
        }

        // 设置返回数据的格式
        w.Header().Set("Content-Type", "application/json")
        // 发送数据
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.PostUserRegist(context.TODO(), &USERINFOSRV.UserRegistRequest{
        //Name: request["name"].(string),
        Mobile:   request["mobile"].(string),
        Password: request["password"].(string),
        SmsCode:  request["sms_code"].(string),
    })

    if err != nil {
        beego.Info("err : ", err)
        http.Error(w, err.Error(), 500)
        return
    }

    if rsp.Errno == utils.RECODE_OK {
        /* // 读取cookie   统一cookie   userlogin
         cookie, err := r.Cookie("userlogin")
         if err == nil || "" == cookie.Value {
             // 创建一个cookie对象
             cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 3600}
             // 对浏览器的cookie进行设置
             http.SetCookie(w, &cookie)
         }*/

        // 用户注册不管之前有没有数据，都重新设置cookie
        // 创建一个cookie对象
        cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 3600}
        // 对浏览器的cookie进行设置
        http.SetCookie(w, &cookie)
        beego.Info("regist set session id ", rsp.SessionId)
    }

    // 准备回传数据
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 设置返回数据的格式
    w.Header().Set("Content-Type", "application/json")
    // 发送给前端
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 用户登录 */
func PostUserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("登陆  PostLogin /api/v1.0/sessions")
    // 接收post发送过来的数据
    var request map[string]interface{}
    // 解析失败，直接返回错误
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    beego.Info(request)

    if request["mobile"].(string) == "" || request["password"].(string) == "" {

        response := map[string]interface{}{
            "errno":  utils.RECODE_DATAERR,
            "errmsg": utils.RecodeText(utils.RECODE_DATAERR),
        }

        // 设置返回数据的格式
        w.Header().Set("Content-Type", "application/json")
        // 发送数据
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }
    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.PostUserLogin(context.TODO(), &USERINFOSRV.UserLoginRequest{
        Mobile:   request["mobile"].(string),
        Password: request["password"].(string),
    })
    if err != nil {
        beego.Info("err : ", err)
        http.Error(w, err.Error(), 500)
        return
    }

    if rsp.Errno == utils.RECODE_OK {

        // 用户登录不管之前有没有数据，都重新设置cookie
        // 创建一个cookie对象
        cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 3600}
        // 对浏览器的cookie进行设置
        http.SetCookie(w, &cookie)

    }
    // 准备回传数据
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 设置返回数据的格式
    w.Header().Set("Content-Type", "application/json")
    // 发送给前端
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 获取Session */
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("获取session信息 GetSession /api/v1.0/session")

    cookie, err := r.Cookie("userlogin")
    if err != nil || cookie == nil || cookie.Value == "" {
        // 直接返回说名用户未登陆
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }
        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    sessionId := cookie.Value
    beego.Info("get session id ", sessionId)
    server := grpc.NewService()
    server.Init()
    //
    //// call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.GetSession(context.TODO(), &USERINFOSRV.SessionRequest{
        //Name: request["name"].(string),
        SessionId: sessionId,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    data := make(map[string]string)
    if utils.RECODE_OK == rsp.Errno {

        data["name"] = rsp.UserName
    }

    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.Errmsg,
        "data":   data,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 获取用户信息 */
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    beego.Info("获取用户信息 GetUserInfo /api/v1.0/user ")

    cookie, err := r.Cookie("userlogin")
    if err != nil || cookie == nil || cookie.Value == "" {
        // 直接返回说名用户未登陆
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }
        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    sessionId := cookie.Value
    beego.Info("sessionid ", sessionId)

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.GetUserInfo(context.TODO(), &USERINFOSRV.UserInfoRequest{
        SessionId: sessionId,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    data := make(map[string]interface{})

    if utils.RECODE_OK == rsp.Errno {

        data["name"] = rsp.Name
        data["user_id"] = rsp.UserId
        data["mobile"] = rsp.Mobile
        data["real_name"] = rsp.RealName // 真实名
        data["id_card"] = rsp.IdCard     // 身份证号
        data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

    }
    // we want to augment the response
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   data,
    }

    beego.Info("data ", data)

    // 设置返回数据的格式
    w.Header().Set("Content-Type", "application/json")
    // 发送给前端
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 用户退出 */
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("DeleteSession  用户退出 /api/v1.0/sessions")

    cookie, err := r.Cookie("userlogin")
    if err != nil || cookie == nil || cookie.Value == "" {
        // 直接返回说名用户未登陆
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }
        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    sessionId := cookie.Value
    beego.Info("sessionid ", sessionId)

    server := grpc.NewService()
    server.Init()

    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.DeleteSession(context.TODO(), &USERINFOSRV.DelSessionRequest{
        //Name: request["name"].(string),
        SessionId: sessionId,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 删除sessionid，用户都退出了，之前的cookie肯定要清除了。

    // 创建一个cookie对象
    ck := http.Cookie{Name: "userlogin", Value: "", Path: "/", MaxAge: -1}
    // 对浏览器的cookie进行设置
    http.SetCookie(w, &ck)

    // 准备回传数据
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 设置返回数据的格式
    w.Header().Set("Content-Type", "application/json")
    // 发送给前端
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 上传用户头像 */
func PostUserAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info("上传头像  PostAvatar /api/v1.0/user/avatar")

    // 获取到前端发送的文件信息

    file, fileHeader, err := r.FormFile("avatar")
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
    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.PostUserAvatar(context.TODO(), &USERINFOSRV.UserAvatarRequest{
        //Name: request["name"].(string),
        SessionId: cookie.Value,
        FileName:  fileHeader.Filename,
        FileSize:  fileHeader.Size,
        Avatar:    fileBuf,
    })
    if err != nil {
        beego.Info("err ", err)
        http.Error(w, err.Error(), 500)
        return
    }

    data := make(map[string]string)
    data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
    // we want to augment the response
    beego.Info("data ", data)

    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   data,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 修改用户名 */
func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info(" 更新用户名 Putuserinfo /api/v1.0/user/name")

    cookie, err := r.Cookie("userlogin")

    if err != nil || cookie == nil || cookie.Value == "" {
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    // 接收前端表单发送内容
    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    beego.Info("request : ", request)
    beego.Info("username : ", request["name"].(string))

    server := grpc.NewService()
    server.Init()
    // call the backend service
    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.PutUserInfo(context.TODO(), &USERINFOSRV.UpdateUserNameRequest{
        SessionId: cookie.Value,
        UserName:  request["name"].(string),
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    data := make(map[string]string)
    data["name"] = rsp.UserName

    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
        "data":   data,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // 返回前端
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

/* 更新实名认证 */
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    beego.Info(" 实名认证 Postuserauth  api/v1.0/user/auth ")

    cookie, err := r.Cookie("userlogin")

    if err != nil || cookie == nil || cookie.Value == "" {
        response := map[string]interface{}{
            "errno":  utils.RECODE_SESSIONERR,
            "errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
        }

        // 回传数据的时候三直接发送过去的并没有设置数据格式
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        return
    }

    var request map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    server := grpc.NewService()
    server.Init()

    exampleClient := USERINFOSRV.NewExampleService("go.micro.srv.UserInfoSrv", server.Client())
    rsp, err := exampleClient.PostUserAuth(context.TODO(), &USERINFOSRV.UpdateUserAuthRequest{
        SessionId: cookie.Value,
        RealName:  request["real_name"].(string),
        IdCard:    request["id_card"].(string),
    })

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 返回给前端的map
    response := map[string]interface{}{
        "errno":  rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 回传数据的时候三直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
