package config

// Logger 日志信息，yaml文件
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      //是否显示行号
	LogInConsole string `yaml:"log_in_console"` //是否打印到控制台
}
