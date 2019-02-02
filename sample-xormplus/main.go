package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/xormplus/xorm"
	"github.com/xormplus/core"
)

type UserInfo struct {  
	Id int
    UserName string
    Age   int
	Desp  string
}


var engine *xorm.Engine

func main() {
    var err error
	
    engine, err = xorm.NewEngine("mysql", "sa:lanshanTest@2018&26g@tcp(221.234.36.70:3307)/test1?charset=utf8mb4")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetMapper(core.GonicMapper{})
	engine.ShowSQL(true)
	
	userInfo := new(UserInfo)
	userInfo.UserName = "myname"
	userInfo.Age = 12
	userInfo.Desp= "第一个用户"
	affected, err := engine.Insert(userInfo)
	if err != nil {
		println(err.Error())
	}else{
		println(affected);
	}
}