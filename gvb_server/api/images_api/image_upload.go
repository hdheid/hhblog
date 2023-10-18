package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"
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
	var resList []image_ser.FileUploadResponse
	for _, file := range fileList {
		//使用函数
		res, fileType, imageHash := service.ServiceApp.ImageService.ImageUploadService(file)
		//如果上传失败
		if !res.IsSuccess {
			resList = append(resList, res)
			continue
		}

		//如果不上传七牛云就储存在本地
		if !global.Config.QiNiu.Enable {
			err = c.SaveUploadedFile(file, res.FileName)
			if err != nil {
				global.Log.Error("图片上传失败：", err)
				res.Msg = "图片上传本地失败!"
				res.IsSuccess = false

				resList = append(resList, res)
				continue
			}
		}

		//将上传成功的图片信息也存储在结构体切片中
		resList = append(resList, res)
		//将图片信息入库
		global.DB.Create(&models.BannerModel{
			Path:      res.FileName,
			Hash:      imageHash,
			Name:      file.Filename,
			ImageType: fileType,
		})
	}

	common.OKWithData(resList, c)
}
