package config

// System 系统配置，yaml文件
type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}
