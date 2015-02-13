package models

type TagProduct struct {
	Id        uint `orm:"fk;auto"`
	ProductId uint `orm:"default(0);index"`
	TagId     uint `orm:"default(0);index"`
}