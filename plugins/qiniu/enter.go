package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_server/config"
	"gvb_server/global"
	"time"
)

/*
这是一个上传七牛云的一个函数，在使用之前，需要下载一个依赖
go get github.com/qiniu/go-sdk/v7
*/

// 获取上传Token
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken
}

// 获取上传的配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}

	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone)) //将q.Zone强转
	cfg.Zone = &zone

	//是否使用https域名
	cfg.UseHTTPS = false

	//上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return cfg
}

func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云上传！")
	}

	qiNiu := global.Config.QiNiu
	if qiNiu.AccessKey == "" || qiNiu.SecretKey == "" {
		return "", errors.New("请配置accessKey以及SecretKey")
	}
	size := float64(len(data)) / 1024 / 1024
	if size > qiNiu.Size {
		return "", errors.New(fmt.Sprintf("图片大小为%.2fMB，超过%dMB，请重新上传", size, global.Config.QiNiu.Size))
	}

	upToken := getToken(qiNiu)
	cfg := getCfg(qiNiu)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		global.Log.Debug("七牛云上传失败！")
		return "", err
	}

	return fmt.Sprintf("%s%s", qiNiu.CDN, key), nil //返回上传的图片的链接,当是公开空间的时候使用这个
}
