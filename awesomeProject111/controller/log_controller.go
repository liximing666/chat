package controller

import (
	"awesomeProject111/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchLog(c *gin.Context) {

	//requestParam := &request.ChannelListRequest{}
	//c.BindJSON(opt)

	channelService := service.NewLogService(c)
	data, _ := channelService.GetAllLog()

	c.JSON(http.StatusOK, data)
	return
}
