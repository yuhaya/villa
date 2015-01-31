package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "villa/routers"
)

func main() {
	orm.Debug = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.Run()
}