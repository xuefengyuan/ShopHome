package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "github.com/garyburd/redigo/redis"
    "github.com/astaxie/beego/orm"
    "ShopHome/HomeWeb/modules"
    "time"
)


// 更新用户名
func (e *Example) PutUserInfo(ctx context.Context, req *example.UpdateUserNameRequest, rsp *example.UpdateUserNameResponse) error {
    //打印被调用的函数
    beego.Info("----- PUT  /api/v1.0/user/name PutUersinfo() --------")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    /* 得到用户发送过来的name */
    beego.Info("name : ",req.UserName)
    beego.Info("name : ",req.SessionId)

    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    /* 从从sessionid获取当前的userid */
    sessionId := req.SessionId
    sessionUserId := bm.Get(utils.GetUserIdKey(sessionId))
    userIdStr, _ := redis.String(sessionUserId, nil)
    userId, _ := strconv.Atoi(userIdStr)

    beego.Info("userid : ", userId)

    o := orm.NewOrm()
    user := models.User{Id: userId, Name: req.UserName}

    num, err := o.Update(&user, "name")
    if err != nil {
        beego.Info("用户名更新失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    beego.Info("index = ",num)

    /* 更新session user_id */

    // n更新用户名
    bm.Put(utils.GetNameKey(sessionId), user.Name, time.Second*3600)
    // 更新session+userid
    bm.Put(utils.GetUserIdKey(sessionId), user.Id, time.Second*3600)

    rsp.UserName = user.Name

    return nil
}
