package models

import (
	"time"
)

type Product struct {
	Id                   uint      `orm:"fk;auto"`
	DestinationId        uint      `orm:"default(0)"`
	PartnerId            uint      `orm:"null;default(0)"`
	CurrencyId           byte      `orm:"default(0)"`
	Name                 string    `orm:"size(30)"`
	NameEn               string    `orm:"null;size(50)"`
	Memo                 string    `orm:"null;size(50)"`
	LocationInfo         string    `orm:"null;size(50)"`
	ShowerRoom           byte      `orm:"default(0)"`
	ShortDescription     string    `orm:"null;size(100)"`
	Description          string    `orm:"null;size(500)"`
	Longitude            float32   `orm:"default(0);digits(7);decimals(4)"` //经度
	Latitude             float32   `orm:"default(0);digits(7);decimals(4)"` //纬度
	AddDate              time.Time `orm:"auto_now_add;type(datetime)"`
	ValidFromDate        time.Time `orm:"auto_now_add;type(date);default(0000-00-00)""`
	ValidToDate          time.Time `orm:"auto_now_add;type(date);default(9999-00-00)""`
	State                byte      `orm:"default(0)"`
	CheckinTime          string    `orm:"default();null"`
	CheckOutTime         string    `orm:"default();null"`
	AdultBedRate         float32   `orm:"default(0);digits(7);decimals(4)"` //每个成人的加床费
	KidBedRate           float32   `orm:"default(0);digits(7);decimals(4)"` //每个儿童加床费
	KidAddBedsMax        byte      `orm:"default(0)"`
	AdultsAddBedsMax     byte      `orm:"default(0)"`
	CancellationClause   string    `orm:"size(255);null"`                   //取消预订条款
	KidClause            string    `orm:"size(255);null"`                   //儿童条款
	LocationTaxRadio     float32   `orm:"default(0);digits(7);decimals(4)"` //地方税率
	ServiceTaxRadio      float32   `orm:"default(0);digits(7);decimals(4)"` //服务费税率
	AirportTransfeRate   float32   `orm:"default(0);digits(7);decimals(4)"` //单程接机费
	Commission           float32   `orm:"default(0);digits(7);decimals(4)"` //提成比例
	DepositRatio         float32   `orm:"default(0);digits(7);decimals(4)"` //面积
	TaxFormula           byte      `orm:"default(0)"`                       //税费计算公式
	Website              string    `orm:"size(100);null"`                   //网址
	ReservationContact   string    `orm:"null;size(20)"`                    //别墅预订联系人
	ReservationTelephone string    `orm:"null;size(20)"`
	ReservationEmail     string    `orm:"null;size(20)"`
}