package models

type InquireCurrency struct {
	Id           uint    `orm:"fk;auto"`
	InquireId    uint    `orm:"index"`
	Name         string  `orm:"size(12)"`
	ExchangeRate float32 `orm:"digits(7);decimals(4)"`
}