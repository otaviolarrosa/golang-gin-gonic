package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/otaviolarrosa/golang-gin-gonic/entity"
	"github.com/otaviolarrosa/golang-gin-gonic/service"
	"github.com/otaviolarrosa/golang-gin-gonic/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	go c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Update(video)
	return nil
}
func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id

	c.service.Delete(video)
	return nil
}
