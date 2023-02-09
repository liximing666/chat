package chat_dao

import "awesomeProject111/dao"

//此文件为数据库的基础文件，所有此数据库的dao均在此文件夹下继承此结构体
type DbBase struct {
	chat_dao.TableBase
}

const DbName = "d_chat"

func (basicDb *DbBase) GetDbName() string {
	return DbName
}
