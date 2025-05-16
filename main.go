package main

import (
	"github.com/anything/smth/1/config"
	"github.com/anything/smth/1/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	r.GET("/:id", controllers.GetDataWithShortedUrl)
	r.POST("/short-url", controllers.GetShortedUrl)

	r.Run(":8080")
}
