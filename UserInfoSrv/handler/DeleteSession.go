package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
)


// 用户退出
func (e *Example) DeleteSession(ctx context.Context, req *example.DelSessionRequest, rsp *example.DelSessionResponse) error {
    beego.Info("DeleteSession  退出登陆 /api/v1.0/session")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    bm, err := RedisInit()

    if err!=nil{
        beego.Info("Redis连接失败",err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    sessionId := req.SessionId
    /* 拼接key  删除session相关信息*/
    // 清除name
    bm.Delete(utils.GetNameKey(sessionId))
    // 清除userId
    bm.Delete(utils.GetUserIdKey(sessionId))
    // 清除手机号
    bm.Delete(utils.GetMobileKey(sessionId))

    return nil
}
