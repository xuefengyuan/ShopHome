package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
    "github.com/garyburd/redigo/redis"
)


// 用户实名
func (e *Example) PostUserAuth(ctx context.Context, req *example.UpdateUserAuthRequest, rsp *example.UpdateUserAuthResponse) error {
    beego.Info(" 实名认证 Postuserauth  api/v1.0/user/auth ")

    //创建返回空间
    rsp.Errno= utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    /* 从从sessionid获取当前的userid */
    sessionId := req.SessionId


    sessionUserId := bm.Get(sessionId + "user_id")
    userIdStr, _ := redis.String(sessionUserId, nil)
    userId, _ := strconv.Atoi(userIdStr)

    beego.Info("userid : ", userId)
    o := orm.NewOrm()
    user := models.User{Id: userId, Real_name: req.RealName, Id_card: req.IdCard}
    index, err := o.Update(&user, "real_name", "id_card")
    if err != nil || index == 0 {
        beego.Info("用户实名信息更新失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // userId
    bm.Put(utils.GetUserIdKey(sessionId), user.Id, time.Second*3600)

    return nil
}
