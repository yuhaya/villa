package models

import (
	"time"
)

type Rate struct {
	Id        uint      `orm:"fk;auto"`
	ProductId uint      `orm:"default(0)"`
	StartDate time.Time `orm:"type(date)"`
	EndDate   time.Time `orm:"type(date)"`
	Bedroom   byte      `orm:"default(0)"`
	MiniStay  byte      `orm:"default(0)"`
}