package bootstrap

import (
	"awesomeProject111/conf"
	"awesomeProject111/dao"
	"awesomeProject111/env"
)

func Init() {
	conf.Init(env.GetEnv())
	chat_dao.Init()
}
