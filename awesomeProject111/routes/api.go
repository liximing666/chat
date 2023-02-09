package routes

import (
	"awesomeProject111/controller"
	"awesomeProject111/model/serializer/request"
	"awesomeProject111/service"
	"github.com/gin-gonic/gin"
	"log"
)

func routeListForApi(r *gin.Engine) *gin.Engine {
	logGroup := r.Group("/log")
	{
		logGroup.POST("/all", controller.SearchLog) //手动接口

	}

	searchGroup := r.Group("/search")
	{
		searchGroup.POST("/get", controller.Chat) //手动接口
	}

	testGroup := r.Group("/test")
	{
		testGroup.POST("/mytest", func(context *gin.Context) {
			params := &request.ChatRequest{}
			context.BindJSON(params)

			log.Println(params.Prompt)

			service := service.NewChatService(context)
			data := service.GetChatRes(params)

			context.JSON(200, data)
		})
	}

	return r
}
