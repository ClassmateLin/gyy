<center><b>gyy</b></center>

### 简介

gyy是一款简洁的go语言web开发框架。



### 安装

`go get github.com/ClassmateLin/gyy`



### 使用

示例:

```go
package main

import (
	"github.com/ClassmateLin/gyy"
	"net/http"
)

type Book struct {
	Id int
	Book string
}


func main()  {
	r := gyy.Default()
	
	
	r.GET("/", func(context *gyy.Context) {
		context.String(http.StatusOK, "Hello gyy")
	})
	
	web := r.Group("web")  // 使用路由组

	web.GET("/", func(context *gyy.Context) {
		context.String(200, "web index") //  返回text/plain
	})
	
	web.POST("/post", func(context *gyy.Context) {
		data := context.Req.PostFormValue("name")
		context.String(200, data)
	})
	
	api := r.Group("api")
	
	api.GET("/book", func(context *gyy.Context) {
		context.JSON(200, Book{1,"Go"})  // 返回json
	})

	r.Run(":8080")
}

```

