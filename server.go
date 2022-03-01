package main

import (
	"github.com/gin-gonic/gin"
	"github.com/otaviolarrosa/golang-gin-gonic/controller"
	"github.com/otaviolarrosa/golang-gin-gonic/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

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
