package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`          // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"输入链接非法" structs:"herf"`       // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"输入图片地址非法" structs:"images"` // 图片
	IsShow bool   `json:"is_show"  msg:"请选择是否展示" structs:"is_show"`                      // 是否展示
}

func main() {
	mp := structs.Map(&AdvertRequest{
		Title:  "标题",
		Href:   "链接",
		Images: "图片",
		IsShow: true,
	})

	fmt.Println(mp)
}
