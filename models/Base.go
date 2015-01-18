package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Base struct {
}

type Admin struct {
	Id   int `orm:"fk;auto"`
	Name string
	Pwd  string
	Date int
}
type Currency struct {
	Id int `orm:"fk;auto"`
}
type Destination struct {
	Id int `orm:"fk;auto"`
}
type Inquire struct {
	Id int `orm:"fk;auto"`
}
type Partner struct {
	Id int `orm:"fk;auto"`
}
type Product struct {
	Id int `orm:"fk;auto"`
}
type Promotion struct {
	Id int `orm:"fk;auto"`
}
type Rate struct {
	Id int `orm:"fk;auto"`
}
type Tag struct {
	Id int `orm:"fk;auto"`
}
type TagProduct struct {
	Id int `orm:"fk;auto"`
}
type User struct {
	Id int `orm:"fk;auto"`
}

func Regist() {
	dbUser := "root"
	dbName := "villa" //数据库名字
	dbPwd := ""
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName) //数据库连接字符串
	orm.RegisterModel(new(Admin), new(Currency), new(Destination), new(Inquire), new(Partner), new(Product), new(Promotion), new(Rate), new(Tag), new(TagProduct), new(User))
	orm.RegisterDriver("mysql", orm.DR_MySQL)        //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", dbLink) //注册数据库，并设置默认数据库
}

func init() {
	Regist()
}
