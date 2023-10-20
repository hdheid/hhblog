# swag的下载和使用

## swag的下载

```go
go install github.com/swaggo/swag/cmd/swag@latest
//当执行过这个之后，其他的项目就只需要执行下面的两条命令即可

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

使用`swag init`可以生成api文档

浏览器访问：`http://ip:hsot/swagger/index.html`

## swag的使用
需要在main函数前面加上这些：
```go
import(
    _ "docs" //或者需要加上一个项目名：_ "项目名/docs"
)

// @title xxx
// @version xxx
// @description xxx
// @host xxx
// @BasePath xxx
func main() {
    
}
```

执行`swag init`之后会出现一个docs的目录

在路由中添加：
```go
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

在每一个api函数前，同样需要添加一些前缀：
```go
// @tags 标签
// @Summary 标题
// @Description 描述,可以有多个
// Param limit query string false "表示单个参数"
// Param data body request.Request   true "表示多个参数"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} gin.H{"msg":"成功"}
func xxxView(c *gin.context){
	
}
```

最后，运行项目，在浏览器中输入 **`127.0.0.1:8080/swagger/index.html`** 即可看到API文档。