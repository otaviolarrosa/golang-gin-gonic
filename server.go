package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otaviolarrosa/golang-gin-gonic/controller"
	"github.com/otaviolarrosa/golang-gin-gonic/middlewares"
	"github.com/otaviolarrosa/golang-gin-gonic/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.Save(c))
	})

	server.Run(":8080")
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
