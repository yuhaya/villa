package models

import (
	"time"
)

type InquireExtra struct {
	Id              uint      `orm:"fk;auto"`
	InquireId       uint      `orm:"index"`
	CardId          string    `orm:"size(30);null"`
	FlightNumber    string    `orm:"size(30);null"` //航班号码
	FlightStartTime time.Time `orm:"null;type(datetime)"`
	FlightEndTime   time.Time `orm:"null;type(datetime)"`
}