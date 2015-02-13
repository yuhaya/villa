package models



type DestinationImg struct {
	Id            uint   `orm:"fk;auto"`
	DestinationId uint   `orm:"index"`
	Img           string `orm:size(120)`
	Type          byte   //图片类型
}