package service

import (
	"awesomeProject111/conf"
	chat_dao "awesomeProject111/dao"
	chat_dao2 "awesomeProject111/dao/chat_dao"
	"github.com/gin-gonic/gin"
	"log"
)

type LogService struct {
	c   *conf.BaseConfig
	d   *chat_dao.Dao
	ctx *gin.Context
}

func NewLogService(ctx *gin.Context) *LogService {
	b := &LogService{
		c:   conf.BaseConf,
		d:   chat_dao.NewDao(ctx),
		ctx: ctx,
	}

	return b
}

/*
	具体的service*******************************************************************************************************
*/

func (t *LogService) GetAllLog() ([]*chat_dao2.SearchLog, error) {
	loggerDao := chat_dao2.NewSearchLogger(t.ctx)
	log.Println(t.ctx.RemoteIP())

	allLog, err := loggerDao.GetAllLog()

	return allLog, err
}
