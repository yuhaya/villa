package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Base struct {
}

func Regist() {
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName") //数据库名字
	dbPwd := beego.AppConfig.String("dbPwd")
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName) //数据库连接字符串
	orm.RegisterModelWithPrefix("villa_",
		new(Admin),
		new(Currency),
		new(DestinationImg),
		new(InquirePriceCost),
		new(InquireVilla),
		new(InquireCurrency),
		new(InquireExtra),
		new(Destination),
		new(Inquire),
		new(Partner),
		new(Product),
		new(ProductImg),
		new(Promotion),
		new(Rate),
		new(Tag),
		new(TagProduct),
		new(User),
	)
	orm.RegisterDriver("mysql", orm.DR_MySQL)        //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", dbLink) //注册数据库，并设置默认数据库
}

func init() {
	Regist()
	createTable()
}

func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}
