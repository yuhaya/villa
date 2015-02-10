package models

import (
	"time"
)

type User struct {
	Id           uint      `orm:"fk;auto"`
	Name         string    `orm:"size(30)"`
	Password     string    `orm:"size(40)"`
	Email        string    `orm:"size(40);null"`
	Phone        string    `orm:"size(40);null"`
	AddDate      time.Time `orm:"auto_now_add;type(datetime)"`
	UserPassport string    `orm:"size(30);null"`
}