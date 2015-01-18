package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var connect_session *mgo.Session
var connect_db *mgo.Database

type BaseModelInterface interface {
}

type BaseModel struct {
}

func (this *BaseModel) GetSession() *(mgo.Session) {

	if connect_session == nil {

		session, err := mgo.Dial(beego.AppConfig.String("mgo_url"))
		connect_session = session
		if err != nil {
			panic(err)
		}
		defer session.Close()
		connect_session = session
		return session
	} else {
		return connect_session
	}

}

func (this *BaseModel) GetDatabase(database string) *(mgo.Database) {
	if database == "" {
		database = beego.AppConfig.String("mgo_database")
	}
	if connect_session == nil {
		this.GetSession()
	}
	connect_db = connect_session.DB(database)
	return connect_db
}

func (this *BaseModel) GetTable(table string) *(mgo.Collection) {
	db_name := ""
	db := this.GetDatabase(db_name)
	table_obj := db.C(table)
	return table_obj
}
