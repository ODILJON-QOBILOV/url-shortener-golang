package main

import (
	"os"
	"log"

	"github.com/anything/smth/1/config"
	"github.com/anything/smth/1/controllers"
	_ "github.com/anything/smth/1/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/joho/godotenv"
)

// @title Url Shortener API
// @version 1.0
// @description url-shortener API

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	config.InitDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.GET("/:id", controllers.GetDataWithShortedUrl)
	r.POST("/short-url", controllers.GetShortedUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
