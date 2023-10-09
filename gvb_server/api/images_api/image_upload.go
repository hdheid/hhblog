package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
	"io/fs"
	"os"
	"path"
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
		filePath := path.Join(global.Config.Upload.Path, file.Filename) //关于图片上传的黑名单和白名单还没有弄
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

		//正常上传
		err := c.SaveUploadedFile(file, filePath)
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
	}

	common.OKWithData(resList, c)
}
