package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
)


// 用户登录
func (e *Example) PostUserLogin(ctx context.Context, req *example.UserLoginRequest, rsp *example.UserLoginResponse) error {
    beego.Info("登陆  PostLogin /api/v1.0/sessions")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    /* 查询数据 */

    // 创建数据库orm句柄
    o := orm.NewOrm()
    // 创建user对象
    var user models.User
    // 创建查询条件句柄
    qs := o.QueryTable("user")
    // 通过qs句柄进行查询
    err := qs.Filter("mobile", req.Mobile).One(&user)

    if err != nil {
        beego.Info("用户登录查询数据失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    /* 密码的校验 */
    if utils.Md5String(req.Password) != user.Password_hash {
        beego.Info("用户登录密码错误")
        rsp.Errno = utils.RECODE_PWDERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 创建sessionid 顺便就把数据返回 */
    sessionId := utils.Md5String(req.Mobile + req.Password)
    rsp.SessionId = sessionId

    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 以sessionid为key的一部分创建session */
    name := user.Name
    if name == "" {
        name = user.Mobile
    }
    beego.Info("name = ",name)
    beego.Info("session id  = ",sessionId)

    // name
    bm.Put(utils.GetNameKey(sessionId), name, time.Second*3600)
    // userId
    bm.Put(utils.GetUserIdKey(sessionId), user.Id, time.Second*3600)
    // 手机号
    bm.Put(utils.GetMobileKey(sessionId), user.Mobile, time.Second*3600)


    return nil
}
