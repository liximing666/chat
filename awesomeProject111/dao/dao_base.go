package chat_dao

import (
	"awesomeProject111/conf"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
	"time"
)

var connections = map[string]*gorm.DB{}

//配置批量插入的大小
const BATCH_SIZE = 500

var DB *gorm.DB

type Dao struct {
	Conf *conf.BaseConfig
	Orm  *gorm.DB
	ctx  *gin.Context
}

type TableModel interface {
	GetDbName() string
	DbInit(table TableModel, ctx *gin.Context)
	GetCtx() *gin.Context
	SetCtx(ctx *gin.Context)
	SetDb(db *gorm.DB)
}

type TableBase struct {
	DB        *gorm.DB              `gorm:"-"`
	ctx       *gin.Context          `gorm:"-"`
	CreatedAt int64                 `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;column:deleted_at" json:"deleted_at"`
}

func (tb *TableBase) DbInit(table TableModel, ctx *gin.Context) {
	dbName := table.GetDbName()
	var db *gorm.DB

	ctxExist := false
	if ctx != nil {
		connInstance, exist := ctx.Get(dbName)
		if exist {
			db = connInstance.(*gorm.DB)
			ctxExist = true
		}
	}

	if !ctxExist {
		//对于脚本等没有经过setdb的中间键的情况，需要重新寻找连接
		if currentConn, exist := GetConn(dbName); exist {
			db = currentConn
		} else {
			panic("找不到" + dbName + "的连接")
		}
	}

	table.SetDb(db)
	table.SetCtx(ctx)
}

func (tb *TableBase) SetCtx(ctx *gin.Context) {
	tb.ctx = ctx
}

func (tb *TableBase) GetCtx() *gin.Context {
	return tb.ctx
}

func (tb *TableBase) GetDbModel() *gorm.DB {
	return tb.DB.Model(tb)
}

func (tb *TableBase) SetDb(db *gorm.DB) {
	tb.DB = db
}

func InitConn(dbName string, prefix string) *gorm.DB {
	dbConf := conf.BaseConf.Mysql

	var mysqlDsn string
	mysqlDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=5s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	//初始化日志
	newLogger := NewLogger(
		Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
			Host:          dbConf.Host,
		},
	)

	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(errors.New("数据库连接错误: " + err.Error()))
	}

	if prefix != "" {
		db.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: true,
		}
	}

	return db
}

func Init() {
	//可以这里定义多个数据库连接
	if conf.BaseConf.Mysql.Host == "" {
		return
	}

	connections["d_chat"] = InitConn("d_chat", "t_")

}

func GetConn(dbName string) (*gorm.DB, bool) {
	db, exist := connections[dbName]
	return db, exist
}

func NewDao(ctx *gin.Context) *Dao {
	d := &Dao{
		Conf: conf.BaseConf,
		ctx:  ctx,
	}
	return d
}
