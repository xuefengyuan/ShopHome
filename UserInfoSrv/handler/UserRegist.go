package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "encoding/json"
    "github.com/astaxie/beego/cache"

    _ "github.com/astaxie/beego/cache/redis"
    _ "github.com/garyburd/redigo/redis"
    _ "github.com/gomodule/redigo/redis"
    "github.com/garyburd/redigo/redis"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
)

type Example struct{}

func RedisInit() (cache.Cache, error) {
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

    return bm, err
}

// 用户注册
func (e *Example) PostUserRegist(ctx context.Context, req *example.UserRegistRequest, rsp *example.UserRegistResponse) error {
    beego.Info("PostRegist  注册 /api/v1.0/users")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 通过手机号获取到短信验证码
    mobile := bm.Get(req.Mobile)

    if mobile == nil {
        beego.Info("验证码获取失败", err)
        rsp.Errno = utils.RECODE_NODATA
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 短信验证码校验
    smsCodeStr, _ := redis.String(mobile, nil)
    if smsCodeStr != req.SmsCode {
        beego.Info("验证码校验失败", err)
        rsp.Errno = utils.RECODE_SMSERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    beego.Info("mobile = ",req.Mobile)
    /* 将数据存入数据库 */
    o := orm.NewOrm()
    user := models.User{Mobile: req.Mobile, Password_hash: utils.Md5String(req.Password), Name: req.Mobile}

    id, err := o.Insert(&user)
    if err != nil {
        beego.Info("用户注册失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    beego.Info("user_id", id)

    /* 创建sessionid  （唯一的随即码）*/
    sessionId := utils.Md5String(req.Mobile + req.Password)
    rsp.SessionId = sessionId

    /* 以sessionid为key的一部分创建session */
    // name 名字暂时使用手机号
    bm.Put(utils.GetNameKey(sessionId), user.Mobile, time.Second*3600)
    // userId
    bm.Put(utils.GetUserIdKey(sessionId), user.Id, time.Second*3600)
    // 手机号
    bm.Put(utils.GetMobileKey(sessionId), user.Mobile, time.Second*3600)
    return nil
}
