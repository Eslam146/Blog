
package main

import (
	"./controller"
	"./middlewares"
	"./repository"
	"./service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)
var(
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService service.VideoService = service.New(videoRepository)
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



	router.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "Input is Valid!"})
		}
	})

	router.DELETE("/videos/:id" , func(ctx *gin.Context) {
		err := videoController.Delete(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "Delete success!"})
		}
	})

	router.PUT("/videos/:id" , func(ctx *gin.Context) {
		err := videoController.Update(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "Update Success"})
		}
	})
	router.Run(":8080")
}
