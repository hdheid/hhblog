package config

// Mysql 数据库信息，yaml文件
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_Level"` //debug是输出全部sql，dev开发环境，release线上环境
}
