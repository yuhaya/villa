package models

import (
	"time"
)

type Inquire struct {
	Id              uint      `orm:"fk;auto"`
	ProductId       uint      `orm:"index"`
	DestinationId   uint      `orm:"index"`
	Uid             uint      `orm:"index"`
	StartDate       time.Time `orm:"type(date)"`
	EndDate         time.Time `orm:"type(date)"`
	Status          byte      `orm:"default(0)"`
	UserIp          string    `orm:"size(20);null"`
	AddDate         time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateDate      time.Time `orm:"auto_now;type(datetime)"`
	UserMessage     string    `orm:"size(250);null"` //用户留言
	Memo            string    `orm:"size(512);null"` //处理备注
	Lock            byte      `orm:"default(0)"`     //上锁 0->不上 1-》上
	Source          byte      `orm:"default(0)"`     //订单来源
	Adults          byte      `orm:"default(0)"`
	Kids            byte      `orm:"default(0)"`
	Bedroom         byte      `orm:"default(0)"`
	KidAddBeds      byte      `orm:"default(0)"`
	AdultsAddBeds   byte      `orm:"default(0)"`
	Invoice         byte      `orm:"default(0)"` //是否开发票
	Coupon          byte      `orm:"default(0)"` //是否有优惠劵
	PayDate         time.Time `orm:"null;type(date)"`
	TaxFormula      byte      `orm:"default(0)"` //税费计算公式
	DeprecatedReson string    `orm:"size(250);null"`
}