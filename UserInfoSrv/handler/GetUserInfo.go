package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "strconv"
    "github.com/garyburd/redigo/redis"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
)


// 获取用户信息
func (e *Example) GetUserInfo(ctx context.Context, req *example.UserInfoRequest, rsp *example.UserInfoResponse) error {

    beego.Info("获取用户信息 GetUserInfo /api/v1.0/user ")

    /* 初始化错误码 */
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg  = utils.RecodeText(rsp.Errno)

    sessionId := req.SessionId

    bm, err := RedisInit()

    if err!=nil{
        beego.Info("Redis连接失败",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    userId := bm.Get(utils.GetUserIdKey(sessionId))
    beego.Info("user id ",userId)

    userIdStr,_ := redis.String(userId,nil)
    beego.Info("user id str ",userIdStr)
    id,_ := strconv.Atoi(userIdStr)
    beego.Info("user id int ",id)

    // 设置用户id,根据用户id查询数据
    user := models.User{Id: id}
    o := orm.NewOrm()
    err = o.Read(&user)

    if err != nil {
        beego.Info("获取用户信息失败",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 将信息返回 */
    rsp.UserId = strconv.Itoa(user.Id)
    rsp.Name = user.Name
    rsp.Mobile = user.Mobile
    rsp.IdCard = user.Id_card
    rsp.RealName = user.Real_name
    rsp.AvatarUrl = user.Avatar_url

    return nil
}
