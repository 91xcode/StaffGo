package common

import (
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

//数据库操作引擎
func NewMysqlEngine() *xorm.Engine{
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/grpc_test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//iris当中我们都讲过 不必多讲
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxOpenConns(10)

	//返回引擎
	return engine
}