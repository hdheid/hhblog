package main

import "gvb_server/models"

func main() {
	models.ArticleModel{}.RemoveIndex()
}
