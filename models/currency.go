package models

import (
	"time"
)

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