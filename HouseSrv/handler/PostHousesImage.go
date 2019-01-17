package handler

import (
    "context"

    example "ShopHome/HouseSrv/proto/example"
    "github.com/astaxie/beego"
    "ShopHome/HomeWeb/utils"
    "path"
    "strconv"
    "ShopHome/HomeWeb/modules"
    "github.com/astaxie/beego/orm"
)

// 上传房屋图片
func (e *Example) PostHousesImage(ctx context.Context, req *example.HousesImageRequest, rsp *example.HousesImageResponse) error {
    beego.Info("发送房屋图片 PostHousesImage  /api/v1.0/houses/:id/images")
    rsp.Errno = utils.RECODE_OK
    rsp.ErrMsg = utils.RecodeText(rsp.Errno)

    size := len(req.Image)
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
    // beego.Info("ext 1 ", ext[1:])
    /* 调用fdfs函数上传到图片服务器 */
    fileId, err := utils.UploadByBuffer(req.Image, ext[1:])
    if err != nil {
        beego.Info("房源图片上传失败")
        rsp.Errno = utils.RECODE_DATAERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 得到fileid */
    beego.Info("file id : ",fileId)

    houseid, _ := strconv.Atoi(req.HouseId)
    house := models.House{Id: houseid}
    o := orm.NewOrm()
    err = o.Read(&house)


    if err != nil {
        beego.Info("房源图片上传，查询房源失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // TODO 这里应该再加上一个用户判断
    beego.Info("huser user id = ",house.User.Id)
    /* 判断index_image_url 是否为空 */
    if house.Index_image_url == "" {
        /* 空就把这张图片设置为主图片 */
        house.Index_image_url = fileId
    }
    /* 将该图片添加到 house的全部图片当中 */
    houseImage := models.HouseImage{House: &house, Url: fileId}

    house.Images = append(house.Images,&houseImage)
    _, err = o.Insert(&houseImage)
    // 将图片对象插入表单之中
    if err != nil {
        beego.Info("房源图片上传，插入图片数据库失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    // 对house表进行更新

    _, err = o.Update(&house)
    if err != nil {
        beego.Info("房源图片上传，更新房源数据库失败")
        rsp.Errno = utils.RECODE_DBERR
        rsp.ErrMsg = utils.RecodeText(rsp.Errno)
        return nil
    }

    /* 返回正确的数据回显给前端 */
    rsp.ImgUrl = fileId

    return nil
}
