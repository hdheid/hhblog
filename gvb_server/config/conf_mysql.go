package config

// Mysql 数据库信息，yaml文件
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Config   string `yaml:"config"` //高级配置，像 charset 等
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //debug是输出全部sql，dev开发环境，release线上环境
}

// Dsn 连接数据库的 url
func (m Mysql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Db + "?" + m.Config
}
