package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

// ImageWhitelist 图片白名单列表
var ImageWhitelist = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"tiff",
	"ico",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// ImageUploadService 专门处理文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse, fileType ctype.ImageType, imageHash string) {
	//定义默认值
	res.FileName = path.Join(global.Config.Upload.Path, file.Filename)
	res.Msg = "上传成功！"
	res.IsSuccess = false
	fileType = ctype.Local //图片储存在哪，默认为本地

	if !IsWhite(file.Filename) { //关于图片上传的白名单
		res.Msg = "非法文件，请重新上传！"
		return
	}

	//判断大小，过大的文件不上传
	size := float64(file.Size) / float64(1024*1024)
	//将过大没有上传的文件信息储存在结构体切片中
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小为%.2fMB，超过%dMB，请重新上传", size, global.Config.Upload.Size)
		return
	}

	fileObj, err := file.Open() //打开文件
	if err != nil {
		global.Log.Error("获取文件对象失败：", err)
	}
	byteData, err := io.ReadAll(fileObj) //读取文件信息
	if err != nil {
		global.Log.Error("读取文件信息失败：", err)
	}
	//哈希值的作用是判断两张图片内容是否是一样的，假如某一张图片只修改了名字，那么这两张图片的哈希值是一样的，就能说明这两张图片其实是一张图片
	//将哈希值存入数据库，就可以很方便地判断这张图片是否已经存在。
	imageHash = utils.Md5(byteData) //转化为哈希值
	//判断数据库中是否存在该图片
	var banner models.BannerModel
	err = global.DB.Take(&banner, "hash = ?", imageHash).Error
	if err == nil {
		//表示找到了该图片
		res.FileName = banner.Path
		res.Msg = "图片已经存在！"
		return
	}

	if global.Config.QiNiu.Enable {
		filePath, err := qiniu.UploadImage(byteData, file.Filename, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error("图片上传七牛云失败：", err)
			res.Msg = err.Error()
			return
		}

		//将上传成功的图片信息存储在结构体切片中
		res.FileName = filePath
		res.Msg = "上传七牛云成功!"

		//修改默认type为七牛
		fileType = ctype.QiNiu
	}

	res.IsSuccess = true //到这里代表上成功，所以为true
	return
}

// IsWhite 判断文件是否在白名单中
func IsWhite(fileNmae string) bool {
	nameList := strings.Split(fileNmae, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1]) //拿到文件的后缀,将其统一变为小写的形式

	global.Log.Debug(suffix)

	for _, str := range ImageWhitelist { //如果该后缀在白名单中，返回真
		if suffix == str {
			return true
		}
	}

	return false
}
