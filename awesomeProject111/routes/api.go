package routes

import (
	"awesomeProject111/controller"
	"awesomeProject111/model/serializer/request"
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
			service := service.NewIpService(context)
			data, _ := service.GetAscriptionPlaceByIp(&request.GetIpRequest{Ip: "202.14.56.4"})

			context.JSON(200, data)
		})
	}

	return r
}
