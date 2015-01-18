package main

import (
	"github.com/astaxie/beego"
	_ "villa/routers"
)

func main() {
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.Run()
}
