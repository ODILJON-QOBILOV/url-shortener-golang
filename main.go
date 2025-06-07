package main

import (
	"github.com/anything/smth/1/config"
	"github.com/anything/smth/1/controllers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/anything/smth/1/docs"
    "github.com/swaggo/files"
)

// @title Url Shortener API
// @version 1.0
// @description url-shortener API

func main() {
	config.InitDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.GET("/:id", controllers.GetDataWithShortedUrl)
	r.POST("/short-url", controllers.GetShortedUrl)

	r.Run(":8080")
}
