package handler

import (
    "context"

    example "ShopHome/UserInfoSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "path"
    "strconv"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
    "github.com/garyburd/redigo/redis"
)


// 上传用户头像
func (e *Example) PostUserAvatar(ctx context.Context, req *example.UserAvatarRequest, rsp *example.UserAvatarResponse) error {
    beego.Info("上传头像  PostAvatar /api/v1.0/user/avatar")

    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    size := len(req.Avatar)
    // 图片数据验证
    if req.FileSize != int64(size) {
        beego.Info("传输数据丢失")
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }


    /* 获取文件的后缀名 */
    ext := path.Ext(req.FileName)
    beego.Info("ext ", ext)
    //beego.Info("ext 1 ", ext[1:])
    /* 调用fdfs函数上传到图片服务器 */
    fileId, err := utils.UploadByBuffer(req.Avatar, ext[1:])
    if err != nil {
        beego.Info("头像上传失败")
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /*得到fileid*/
    beego.Info(fileId)
    sessionId := req.SessionId
    bm, err := RedisInit()
    if err != nil {
        beego.Info("Redis连接失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    sessionUserId := bm.Get(sessionId + "user_id")
    userIdStr, _ := redis.String(sessionUserId,nil)
    userId, _ := strconv.Atoi(userIdStr)
    /* 将图片的存储地址（fileid）更新到user表中 */
    user := models.User{Id: userId, Avatar_url: fileId}
    o := orm.NewOrm()
    _, err = o.Update(&user, "avatar_url")

    if err!= nil {
        beego.Info("用户信息数据库更新失败", err)
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }
    /* 回传fielid */
    rsp.AvatarUrl = fileId
    return nil
}
