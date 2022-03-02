package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otaviolarrosa/golang-gin-gonic/controller"
	"github.com/otaviolarrosa/golang-gin-gonic/middlewares"
	"github.com/otaviolarrosa/golang-gin-gonic/repository"
	"github.com/otaviolarrosa/golang-gin-gonic/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
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
		err := VideoController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Video input is Valid!!"})
		}
	})

	server.PUT("/videos/:id", func(c *gin.Context) {
		err := VideoController.Update(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Sucess"})
		}
	})

	server.DELETE("/videos/:id", func(c *gin.Context) {
		err := VideoController.Delete(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Sucess"})
		}
	})

	server.Run(":8080")
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
