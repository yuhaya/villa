package models

type Tag struct {
	Id   uint   `orm:"fk;auto"`
	Name string `orm:"size(30)"`
	Type byte   `orm:"default(0)"` //tag类型，1-风格， 2-位置，3-主题，特色 , 4-情景，5-免费服务项目，6-基础设施，7-活动设施，8-综合设施，9-周边
}