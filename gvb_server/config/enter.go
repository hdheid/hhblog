package config

// Config 将yaml文件的配置映射到结构体中去
type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
