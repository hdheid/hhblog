package main

import (
	"fmt"
	"gvb_server/utils"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(utils.HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(utils.CheckPwd("$2a$04$uGsxzc.JM7P6eoDEvM0qhuh7ECyA5FCeA7RSTinKPerrz0Gl.JYK6", "123456"))
}
