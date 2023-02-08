package routes

import (
	"awesomeProject111/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetRoutes(r *gin.Engine) *gin.Engine {
	r = routeListForApi(r) //API路由定义

	return r
}

func New() *http.Server {
	// init static file handler
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r = SetRoutes(r) //路由定义

	server := &http.Server{
		Addr:         conf.BaseConf.HttpServerConfig.Addr,
		Handler:      r,
		ReadTimeout:  time.Duration(conf.BaseConf.HttpServerConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.BaseConf.HttpServerConfig.WriteTimeout) * time.Second,
	}

	return server
}
