
package main

import (
	"./controller"
	"./service"
	"github.com/gin-gonic/gin"
	"net/http"
)
var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main() {
	router := gin.Default()

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
	router.Run(":8080")
}
