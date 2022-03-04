package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hello_peer_api_spec/docs"
)

// @title Hello Peer API
// @version 0.1
// @description Hello Peer是一款基于兴趣的社交应用。
// @termsOfService API文档仅用于Hello Peer研发使用。

// @contact.name Frank Sheng
// @contact.email 726569998@qq.com

// @host 1.12.243.73:8686
// @BasePath /api/v0.1
func main() {
	r := gin.Default()
	url := ginSwagger.URL("http://localhost:8181/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8181")
}
