package main


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	db *xorm.EngineGroup
)

func init() {
	conns := []string{"root:@tcp(127.0.0.1:3306)/go_demos"}

	var err error
	db, err = xorm.NewEngineGroup("mysql", conns)
	if err != nil {
		panic(err.Error())
	}

	db.ShowSQL(true)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
}
