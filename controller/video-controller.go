package controller
import (
	 "../entity"
	 "../service"
	 "github.com/gin-gonic/gin"
)
type VideoController interface {
	Save(ctx *gin.Context ) error
	FindAll() []entity.Video
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
