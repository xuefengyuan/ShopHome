package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "github.com/garyburd/redigo/redis"
)

// 获取Session信息
func (e *Example) GetSession(ctx context.Context, req *example.SessionRequest, rsp *example.SessionResponse) error {
    beego.Info("获取session信息 GetSession /api/v1.0/session")

    rsp.Errno = utils.RECODE_OK
    rsp.Errmsg = utils.RecodeText(rsp.Errno)

    bm, err := RedisInit()
    if err != nil {
        beego.Info("redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.Errmsg = utils.RecodeText(rsp.Errno)
    }

    beego.Info("session id = ",req.SessionId)

    /* 没有缓存返回失败 */
    userName := bm.Get(utils.GetNameKey(req.SessionId))
    if userName == nil {
        beego.Info("用户信息不存在", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.Errmsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    beego.Info("username",userName)
    /* 有就返回成功 */
    rsp.UserName, _ = redis.String(userName, nil)

    return nil
}