package handler

import (
    "context"

    _"github.com/astaxie/beego/cache/redis"
    _"github.com/garyburd/redigo/redis"
    _"github.com/gomodule/redigo/redis"
    example "ShopHome/VerifyCodeSrv/proto/example"
    "github.com/astaxie/beego"
    "github.com/afocus/captcha"
    "image/color"
    "ShopHome/HomeWeb/utils"
    "encoding/json"
    "github.com/astaxie/beego/cache"
    "time"
    "reflect"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "github.com/garyburd/redigo/redis"
    "math/rand"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetImageCode(ctx context.Context, req *example.ImageCodeRequest, rsp *example.ImageCodeResponse) error {
    beego.Info("获取验证码图片 GetImageCode /api/v1.0/imagecode/:uuid")

    /*生成验证码图片*/
    // 创建图片句柄
    cap := captcha.New()

    // 这个位置有坑，要注意字体文件的位置
    if err := cap.SetFont("comic.ttf"); err != nil {
        panic(err.Error())
    }

    // 设置图片大小
    cap.SetSize(91, 41)
    // 设置干扰强度
    cap.SetDisturbance(captcha.NORMAL)
    // 设置前景色
    cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
    // 设置背景色
    cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

    // 生存随即的验证码图片
    img, str := cap.Create(4, captcha.NUM)
    /*将uuid和随即验证码进行缓存*/
    // 配置缓存参数
    redisConf := map[string]string{
        "key":   utils.G_server_name,
        "conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
        "dbNum": utils.G_redis_dbnum,
    }
    beego.Info(redisConf)
    // 将map进行转化成为json
    redisConfJs, _ := json.Marshal(redisConf)
    // 创建redis句柄
    bm, err := cache.NewCache("redis", string(redisConfJs))

    if err != nil {
        beego.Info("redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
    }
    // 验证码与uuid进行缓存
    bm.Put(req.Uuid, str, time.Second*300)

    // 图片解引用
    img1 := *img
    img2 := *img1.RGBA

    // 返回错误信息
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    // 返回图片拆分
    rsp.Pix = []byte(img2.Pix)
    rsp.Stride = int64(img2.Stride)

    rsp.Max = &example.ImageCodeResponse_Point{X: int64(img2.Rect.Max.X), Y: int64(img2.Rect.Max.Y)}
    rsp.Min = &example.ImageCodeResponse_Point{X: int64(img2.Rect.Min.X), Y: int64(img2.Rect.Min.Y)}

    return nil
}

func (e *Example) GetSmsCode(ctx context.Context, req *example.SmsCodeRequest, rsp *example.SmsCodeResponse) error {
    beego.Info("获取短信验证码 GetSmsCode api/v1.0/smscode/:mobile ")

    //初始化返回值
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    /* 1、验证手机号是否存在 */

    // 创建数据库orm句柄
    o := orm.NewOrm()
    // 使用手机号作为查询条件
    user := models.User{Mobile: req.Mobile}

    err := o.Read(&user)
    // 如果不报错就说明查找到了
    // 查找到就说明手机号存在
    if err == nil {
        beego.Info("用户已存在")
        rsp.Errno = utils.RECODE_MOBILEERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 2、验证图片验证码是否正确 */
    // 连接Redis,配置缓存参数
    redisConf := map[string]string{
        "key": utils.G_server_name,
        //127.0.0.1:6379
        "conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
        "dbNum": utils.G_redis_dbnum,
    }
    beego.Info(redisConf)

    // 将map进行转化成为json
    redisConfJs, _ := json.Marshal(redisConf)

    // 创建redis句柄
    bm, err := cache.NewCache("redis", string(redisConfJs))
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 通过uuid查找图片验证码的值进行对比
    imageCode := bm.Get(req.Uuid)
    if imageCode == nil {
        beego.Info("Redis 获取图片验证码失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    // reflect.TypeOf(value)会返回当前数据的变量类型
    beego.Info(reflect.TypeOf(imageCode), imageCode)
    // 通过Redis格式转换
    imageCodeStr, _ := redis.String(imageCode, nil)

    if imageCodeStr != req.ImageCode {
        beego.Info("图片验证码错误...", err)
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 3、调用 短信接口发送短信 */

    // 创建随机数
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    code := r.Intn(9999) + 1001
    beego.Info("验证码：", code)

    //真正调用发短信的就注释掉了,需要替换成你自己的appid和appkey，或者调用别的平台的，功能是完整的
/*
    // 短信map
    messageconfig := make(map[string]string)
    // id
    messageconfig["appid"] = "xxxx"
    // key
    messageconfig["appkey"] = "xxxxxxxxxxx"
    // 加密方式
    messageconfig["signtype"] = "md5"

    //短信操作对象
    messagexsend := submail.CreateMessageXSend()
    // 短信发送到那个手机号
    submail.MessageXSendAddTo(messagexsend, req.Mobile)
    // 短信发送的模板
    submail.MessageXSendSetProject(messagexsend, "CNrgX")
    // 发送的验证码
    submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(code))
    submail.MessageXSendAddVar(messagexsend, "time", "3分钟")
    // 发送
    fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))
*/


    /* 4、将短信验证码存入缓存数据库 */
    err = bm.Put(req.Mobile, code, time.Second*300)
    if err != nil {
        beego.Info("Redis 保存失败",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    return nil
}
