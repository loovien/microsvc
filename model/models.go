package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	x   *xorm.Engine
	err error
)

func init() {
	dns := "live:admin@tcp(192.168.163.184)/test?charset=utf8&parseTime=true"
	x, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		panic(err)
	}
}
