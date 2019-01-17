package utils

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/weilaihui/fdfs_client"
    "fmt"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求

func AddDomain2Url(url string) (domain_url string) {
    domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

    return domain_url
}

/* Md5加密*/
func Md5String(s string)string  {
    // 创建1个md5对象
    md := md5.New()
    md.Write([]byte(s))
    return hex.EncodeToString(md.Sum(nil))
}

/* 上传二进制文件到fdfs */
func UploadByBuffer(fileBuffer []byte,fileExt string) (fileId string,err error) {
    // 创建连接句柄，配置文件要写全路径
    fdfsClient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
    if err != nil {
        fmt.Println("创建fdfs句柄失败,",err)
        fileId = ""
        return
    }



    fdReq, err := fdfsClient.UploadByBuffer(fileBuffer, fileExt)
    if err != nil {
        fmt.Println("文件上传失败")
        fileId = ""
    }

    fmt.Println(fdReq.GroupName)
    fmt.Println(fdReq.RemoteFileId)
    return fdReq.RemoteFileId,nil
    //fdfsClient.DeleteFile(fdReq.RemoteFileId)
}