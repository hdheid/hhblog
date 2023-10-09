package config

/*
关于照片上传本地的一些配置，像上传图片大小限制，还有上传图片的路径等
*/

type Upload struct {
	Size int    `yaml:"size" json:"size"` //图片上传的大小
	Path string `yaml:"path" json:"path"` //图片上传的目录
}
