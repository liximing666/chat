package routes

import (
	"awesomeProject111/controller"
	"awesomeProject111/service"
	"github.com/gin-gonic/gin"
)

func routeListForApi(r *gin.Engine) *gin.Engine {
	logGroup := r.Group("/log")
	{
		logGroup.POST("/all", controller.SearchLog) //手动接口

	}

	testGroup := r.Group("/test")
	{
		testGroup.GET("/mytest", func(context *gin.Context) {
			service := service.NewChatService(context)
			data := service.GetChatRes("hahah")

			context.JSON(200, data)
		})
	}

	return r
}
