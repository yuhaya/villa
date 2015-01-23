package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego"
	"time"
)

type Base struct {
}

//管理员表
type Admin struct {
	Id   uint      `orm:"fk;auto"`
	Name string    `orm:"size(20)"`
	Pwd  string    `orm:"size(40)"`
	Date time.Time `orm:"auto_now_add"`
}

//当前币种对人民币的汇率
type Currency struct {
	Id           uint      `orm:"fk;auto"`
	Name         string    `orm:"size(12)"`
	ExchangeRate float32   `orm:"digits(7);decimals(4)"`
	UpdateTime   time.Time `orm:"auto_now"`
}

func (c *Currency) TableEngine() string {
	return "MYISAM"
}

//目的地
type Destination struct {
	Id         uint   `orm:"fk;auto"`
	CurrencyId uint   `orm:"index"`
	Name       string `orm:"size(30);null,default()"`
	NameEn     string `orm:"size(60);null,default()"`
	Level      byte   `orm:"index"`
	Left       uint16
	Right      uint16
	SortNum    uint16 `orm:"default(0)"` //排序
	Memo       string `orm:"size(256)"`
	Tag        byte   `orm:"default(0)"` //标签 0=>默认
	Status     byte   `orm:"default(1)"` //是否可用 1=>可用 0=>不可用
	Show       byte   `orm:"default(0)"` //是否显示 1=>显示 0=>不显示
}

type DestinationImg struct {
	Id            uint   `orm:"fk;auto"`
	DestinationId uint   `orm:"index"`
	Img           string `orm:size(120)`
	Type          byte   //图片类型
}

type Inquire struct {
	Id uint `orm:"fk;auto"`
}
type Partner struct {
	Id uint `orm:"fk;auto"`
}
type Product struct {
	Id uint `orm:"fk;auto"`
}
type Promotion struct {
	Id uint `orm:"fk;auto"`
}
type Rate struct {
	Id uint `orm:"fk;auto"`
}
type Tag struct {
	Id uint `orm:"fk;auto"`
}
type TagProduct struct {
	Id uint `orm:"fk;auto"`
}
type User struct {
	Id uint `orm:"fk;auto"`
}

func Regist() {
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName") //数据库名字
	dbPwd := beego.AppConfig.String("dbPwd")
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName) //数据库连接字符串
	orm.RegisterModelWithPrefix("villa_", new(Admin), new(Currency), new(Destination), new(Inquire), new(Partner), new(Product), new(Promotion), new(Rate), new(Tag), new(TagProduct), new(User))
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
