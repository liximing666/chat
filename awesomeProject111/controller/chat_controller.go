package controller

import (
	"awesomeProject111/model/serializer/request"
	"awesomeProject111/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Chat(c *gin.Context) {

	requestParam := &request.ChatRequest{}
	c.BindJSON(requestParam)

	chatService := service.NewChatService(c)
	data := chatService.GetChatRes(requestParam)

	c.JSON(http.StatusOK, data)
	return
}
