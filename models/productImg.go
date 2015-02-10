package models

type ProductImg struct {
	Id        uint   `orm:"fk;auto"`
	ProductId uint   `orm:"default(0)"`
	Name      string `orm:"size(30)"`
}