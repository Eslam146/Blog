package controller

import (
	"../entity"
	"../service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type VideoController interface {
	Save(ctx *gin.Context ) error
	FindAll() []entity.Video
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &controller{ service: s}
}

func (c *controller )Save(ctx *gin.Context ) error{
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil{
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) FindAll() []entity.Video{
	return c.service.FindAll()
}

func  (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil{
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"),0,0)
	video.ID = id
	c.service.Delete(video)
	return nil
}

func  (c *controller) Update(ctx *gin.Context) error{
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil{
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"),0,0)
	video.ID = id
	c.service.Update(video)
	return nil
}