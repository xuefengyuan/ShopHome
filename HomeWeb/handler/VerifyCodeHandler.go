package handler

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    VERIFYCODESRV "ShopHome/VerifyCodeSrv/proto/example"
    "github.com/astaxie/beego"
    "context"
    "github.com/micro/go-grpc"
    "image"
    "image/png"
    "github.com/afocus/captcha"
    "regexp"
    "ShopHome/HomeWeb/utils"
)

/* 获取图片验证码 */
func GetImageCode(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
    beego.Info("获取验证码图片 GetImageCode /api/v1.0/imagecode/:uuid")


    server := grpc.NewService()
    server.Init()
    uuid := ps.ByName("uuid")

    beego.Info("uuid : ",uuid)

    exampleClient := VERIFYCODESRV.NewExampleService("go.micro.srv.VerifyCodeSrv", server.Client())
    rsp, err := exampleClient.GetImageCode(context.TODO(), &VERIFYCODESRV.ImageCodeRequest{
        //Name: request["name"].(string),
        Uuid:uuid,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }


    // 接收图片信息的 图片格式
    var img image.RGBA

    img.Stride = int(rsp.Stride)
    img.Pix = []uint8(rsp.Pix)

    img.Rect.Min.X = int(rsp.Min.X)
    img.Rect.Min.Y = int(rsp.Min.Y)

    img.Rect.Max.X = int(rsp.Max.X)
    img.Rect.Max.Y = int(rsp.Max.Y)

    var image captcha.Image
    image.RGBA = &img

    // 将图片发送给浏览器
    png.Encode(w, image)

}

/* 获取短信验证码 */
func GetSmsCode(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
    beego.Info("获取短信验证码 GetSmsCode api/v1.0/smscode/:mobile ")


    // 通过传入参数URL下Query获取前端的在url里的带参
    beego.Info(r.URL.Query())

    text := r.URL.Query()["text"][0]
    uuid := r.URL.Query()["id"][0]
    mobile := ps.ByName("mobile")


    // 1、通过正则进行手机号的判断
    mobileReg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
    // 通过条件判断字符串是否匹配规则 返回正确或失败
    isMobile := mobileReg.MatchString(mobile)
    // 如果手机号不匹配，那就直接返回错误不调用服务
    if !isMobile {
        // 创建返回数据的map
        response := map[string]interface{}{
            "errno":  utils.RECODE_MOBILEERR,
            "errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
        }
        // 设置返回数据的格式
        w.Header().Set("Content-Type", "application/json")
        // 发送数据
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
    }

    server := grpc.NewService()
    server.Init()

    exampleClient := VERIFYCODESRV.NewExampleService("go.micro.srv.VerifyCodeSrv", server.Client())
    rsp, err := exampleClient.GetSmsCode(context.TODO(), &VERIFYCODESRV.SmsCodeRequest{
        ImageCode:text,
        Uuid:uuid,
        Mobile:mobile,
    })

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 创建返回数据的map
    response := map[string]interface{}{
        "errno": rsp.Errno,
        "errmsg": rsp.ErrMsg,
    }

    // 回传数据的时候直接发送过去的并没有设置数据格式
    w.Header().Set("Content-Type", "application/json")
    // encode and write the response as json
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
