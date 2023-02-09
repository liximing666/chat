package service

import (
	"awesomeProject111/conf"
	chat_dao "awesomeProject111/dao"
	chat_dao2 "awesomeProject111/dao/chat_dao"
	"awesomeProject111/model/serializer/request"
	"awesomeProject111/model/serializer/response"
	"awesomeProject111/pkg/ecode"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func (t *ChatService) GetChatRes(requestParams *request.ChatRequest) response.ChatResponse {
	data := response.ChatResponse{}

	//归属地屏蔽
	ip := t.ctx.RemoteIP()

	ipService := NewIpService(t.ctx)
	ipRes, code := ipService.GetAscriptionPlaceByIp(ip)
	if code != ecode.OK {
		log.Println(code.Message())
		return data
	}

	if strings.ContainsAny(conf.BaseConf.Black, ipRes.SuccessData.Result.Detailed) {
		return data
	}

	req := NewChatReq(requestParams.Prompt)
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
		log.Println(err.Error())

	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}
	log.Println(string(respBytes))
	json.Unmarshal(respBytes, &data)

	//访问记录
	go func() {
		logDao := chat_dao2.NewSearchLogger(t.ctx)
		err = logDao.CreateSearchLog(&chat_dao2.SearchLog{
			Ip:      ip,
			Content: requestParams.Prompt,
		})
		if err != nil {
			log.Println(err.Error())
		}
	}()

	return data
}
