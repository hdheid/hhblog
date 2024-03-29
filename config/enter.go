package config

// Config 将yaml文件的配置映射到结构体中去
type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Jwy      Jwy      `yaml:"jwy"`
	Email    Email    `yaml:"email"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
	TxMap    TxMap    `yaml:"tx_map"`
	ES       ES       `yaml:"es"`
}
