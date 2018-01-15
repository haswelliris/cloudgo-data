package entities

import (
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	// register go-sql driver
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:test@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	engine.TZLocation, err = time.LoadLocation("Asia/Shanghai")
	checkErr(err)
	engine.SetMapper(core.SameMapper{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
