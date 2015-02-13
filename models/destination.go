package models

//目的地
type Destination struct {
	Id         uint   `orm:"fk;auto"`
	CurrencyId uint   `orm:"index"`
	Name       string `orm:"size(30);null;default()"`
	NameEn     string `orm:"size(60);null;default()"`
	Level      byte   `orm:"index"`
	Left       uint16
	Right      uint16
	SortNum    uint16 `orm:"default(0)"` //排序
	Memo       string `orm:"size(256)"`
	Tag        byte   `orm:"default(0)" valid:"Min(0)`      //标签 0=>默认
	Status     byte   `orm:"default(1)" valid:"Range(0, 1)` //是否可用 1=>可用 0=>不可用
	Show       byte   `orm:"default(0)" valid:"Range(0, 1)` //是否显示 1=>显示 0=>不显示
}