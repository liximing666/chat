package service

import (
	"awesomeProject111/conf"
	chat_dao "awesomeProject111/dao"
	"awesomeProject111/model/serializer/response"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ChatService struct {
	c   *conf.BaseConfig
	d   *chat_dao.Dao
	ctx *gin.Context
}

func NewChatService(ctx *gin.Context) *ChatService {
	b := &ChatService{
		c:   conf.BaseConf,
		d:   chat_dao.NewDao(ctx),
		ctx: ctx,
	}

	return b
}

/*
	具体的service*******************************************************************************************************
*/

type ChatReq struct {
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	Model            string  `json:"model"`
}

func NewChatReq(prompt string) *ChatReq {
	return &ChatReq{
		Prompt:           prompt,
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Model:            "text-davinci-003",
	}
}

func (t *ChatService) GetChatRes(content string) response.ChatResponse {
	req := NewChatReq(content)
	postBody, _ := json.Marshal(req)

	buffer := bytes.NewBuffer(postBody)
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
	}

	request.Header.Set("Content-Type", "application/json")                           //添加请求头
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.BaseConf.Key)) //添加请求头
	client := http.Client{}                                                          //创建客户端
	resp, err := client.Do(request.WithContext(context.TODO()))                      //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)

	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}

	data := response.ChatResponse{}
	json.Unmarshal(respBytes, &data)

	return data
}
