package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

/*
图片上传的黑名单、白名单
黑名单：如果文件后缀名与黑名单中的后缀名重合，就拒绝上传
白名单：只能上传在白名单中出现的后缀
*/

var (
	// ImageWhitelist 图片白名单列表
	ImageWhitelist = []string{
		"jpg",
		"png",
		"jpeg",
		"gif",
		"tiff",
		"ico",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// ImageUploadView 上传单个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Debug("获取上传的文件失败，", err)
		common.FailWithMessage("获取上传的文件失败："+err.Error(), c)
		return
	}

	//设置key为image的时候，无论传递了n个文件，都可以通过这个map读取出来，是一个切片类型
	fileList, ok := form.File["image"]
	if !ok {
		global.Log.Debug("不存在该文件")
		common.FailWithMessage("不存在该文件", c)
		return
	}

	//判断路径是否存在
	//通过调用 os.ReadDir(basePath) 可以获取目录下的所有文件和子目录的名称、大小、权限等信息。您可以进一步遍历该切片，对每个文件或子目录进行相应的操作，比如打印名称、筛选特定类型的文件等。
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	//不存在就创建一个路径
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm) //basePath如果有该路径，则直接返回，没有则递归的创建目录；fs.ModePerm表示创建的目录具有读写权限
		if err != nil {
			global.Log.Error("目录创建失败失败：", err)
		}
	}

	//循环遍历文件
	var resList []FileUploadResponse
	for _, file := range fileList {
		if !IsWhite(file.Filename) { //关于图片上传的白名单
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				Msg:       "非法文件，请重新上传！",
				IsSuccess: false,
			})
			continue
		}

		filePath := path.Join(global.Config.Upload.Path, file.Filename)
		//判断大小，过大的文件不上传
		size := float64(file.Size) / float64(1024*1024)
		//将过大没有上传的文件信息储存在结构体切片中
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				Msg:       fmt.Sprintf("图片大小为%.2fMB，超过%dMB，请重新上传", size, global.Config.Upload.Size),
				IsSuccess: false,
			})
			continue
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
		imageHash := utils.Md5(byteData) //转化为哈希值
		//判断数据库中是否存在该图片
		var banner models.BannerModel
		err = global.DB.Take(&banner, "hash = ?", imageHash).Error
		if err == nil {
			//表示找到了该图片
			resList = append(resList, FileUploadResponse{
				FileName:  banner.Path,
				Msg:       "图片已经存在！",
				IsSuccess: false,
			})
			continue
		}

		//正常上传
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error("图片上传失败：", err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				Msg:       "上传失败！",
				IsSuccess: false,
			})
			continue
		}

		//将上传成功的图片信息也存储在结构体切片中
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			Msg:       "上传成功！",
			IsSuccess: true,
		})

		//将图片信息入库
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: imageHash,
			Name: file.Filename,
		})
	}

	common.OKWithData(resList, c)
}

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
