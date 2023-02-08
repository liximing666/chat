package routes

import (
	"awesomeProject111/controller"
	"github.com/gin-gonic/gin"
)

func routeListForApi(r *gin.Engine) *gin.Engine {
	logGroup := r.Group("/log")
	{
		logGroup.POST("/all", controller.SearchLog) //手动接口

	}

	return r
}
