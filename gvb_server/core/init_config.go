package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

/*
读取yaml文件需要去添加一个yaml.v2的依赖

go get gopkg.in/yaml.v2
*/

/*
InitConf 读取yaml文件的配置
流程为使用 ioutil.ReadFile 将yaml文件数据读取出来，然后使用 yaml.Unmarshal 映射到实例化的 c 中
然后涉及到一些日志打印之类的
*/
func InitConf() {
	const ConfigFile = "settings.yaml"             //yaml文件的路径
	yamlConfig, err := ioutil.ReadFile(ConfigFile) //将yaml文件数据读取出来
	if err != nil {
		panic(fmt.Errorf("获取yaml文件异常：%s", err))
	}

	conf := &config.Config{}               //初始化一个实例
	err = yaml.Unmarshal(yamlConfig, conf) //反序列化,将数据映射到实例化的 c 中
	if err != nil {
		panic(fmt.Errorf("映射到实例化 Config 异常：%s", err))
	}

	log.Println("yaml文件初始化成功！")
	fmt.Println(conf) //将映射好的实例打印到控制台

	global.Config = conf //将conf设置为全局的变量
}
