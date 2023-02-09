package chat_dao

import "github.com/gin-gonic/gin"

type SearchLog struct {
	Id      int    `json:"id"`
	Ip      string `json:"ip"`
	Content string `json:"content"`
	DbBase
}

func NewSearchLogger(ctx *gin.Context) *SearchLog {
	c := &SearchLog{}

	//依赖注入ctx db
	c.DbInit(c, ctx)

	return c
}

/*
	dao API ************************************************************************************************************
*/

func (s *SearchLog) GetAllLog() ([]*SearchLog, error) {
	res := make([]*SearchLog, 0)

	err := s.DB.Model(&SearchLog{}).Find(&res).Error

	return res, err
}

func (s *SearchLog) CreateSearchLog(insertObj *SearchLog) error {
	err := s.DB.Model(&SearchLog{}).Create(insertObj).Error
	return err
}
