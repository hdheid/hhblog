package config

import "fmt"

type ES struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

func (es ES) URL() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
