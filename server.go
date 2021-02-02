
package main

import (
	"./controller"
	"./middlewares"
	"./service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)
var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput()  {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f , os.Stdout)
}

func main() {
	setupLogOutput()
	router := gin.New()
	router.Use(gin.Recovery() , middlewares.Logger() , middlewares.BasicAuth())

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200,videoController.FindAll())
	})

	router.GET("/author/:autherName", func(ctx *gin.Context) {
		autherName := ctx.Param("autherName")
		ctx.JSON(http.StatusOK, videoController.FindByAuthor(autherName) )
	})

	router.GET("/videos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, videoController.FindById(id) )
	})

	router.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "Input is Valid!"})
		}
	})

	router.DELETE("/videos/:id" , func(ctx *gin.Context) {
		id := ctx.Param("id")
		videoController.Delete(id)
		ctx.JSON(http.StatusOK , "item has been removed")
	})

	router.Run(":8080")
}
