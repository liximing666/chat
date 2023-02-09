package service

import (
	"awesomeProject111/conf"
	chat_dao "awesomeProject111/dao"
	"awesomeProject111/model/serializer"
	"awesomeProject111/model/serializer/response"
	"awesomeProject111/pkg/ecode"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

type IpService struct {
	c   *conf.BaseConfig
	d   *chat_dao.Dao
	ctx *gin.Context
}

func NewIpService(ctx *gin.Context) *IpService {
	b := &IpService{
		c:   conf.BaseConf,
		d:   chat_dao.NewDao(ctx),
		ctx: ctx,
	}

	return b
}

/*
	具体的service*******************************************************************************************************
*/

func (t *IpService) getIpRequestUrl(ip string) string {
	return fmt.Sprintf("http://api.k780.com/?app=ip.get&ip=%s&appkey=%s&sign=%s&format=json", ip, serializer.APP_KEY, conf.BaseConf.Sign)
}

//查询归属地
func (t *IpService) GetAscriptionPlaceByIp(ip string) (response.IpResponse, ecode.Code) {
	res := response.IpResponse{}
	url := t.getIpRequestUrl(ip)

	resp, err := http.Get(url)
	if err != nil {
		return res, ecode.INVALID_PARAM.SetMessage(err.Error())
	}

	//todo 请求待二次封装
	respBody, _ := ioutil.ReadAll(resp.Body)
	code := gjson.GetBytes(respBody, "success").Int()

	log.Println(code)

	if code == serializer.Success {
		data := &serializer.IpSuccessData{}
		json.Unmarshal(respBody, &data)
		res.SuccessData = data
	} else {
		res.FailData = &serializer.IpFailData{
			Success: "0",
			Msgid:   "error",
			Msg:     "error",
		}
	}

	return res, ecode.OK
}
