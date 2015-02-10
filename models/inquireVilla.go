package models

type InquireVilla struct {
	Id                 uint    `orm:"fk;auto"`
	InquireId          uint    `orm:"index"`
	UserPassport       string  `orm:"size(30);null"`
	CardId             string  `orm:"size(30);null"`
	BasicFacility      string  `orm:"size(255);null"`                   //基础设施
	IntegratedFacility string  `orm:"size(255);null"`                   //综合设施
	ActionFacility     string  `orm:"size(255);null"`                   //活动设施
	NearbyFacility     string  `orm:"size(255);null"`                   //周边设施
	Tags               string  `orm:"size(255);null"`                   //风格
	ServiceFree        string  `orm:"size(255);null"`                   //免费服务
	LocationTaxRadio   float32 `orm:"default(0);digits(7);decimals(4)"` //地方税率
	ServiceTaxRadio    float32 `orm:"default(0);digits(7);decimals(4)"` //服务费税率
	CancellationClause string  `orm:"size(255);null"`                   //取消预订条款
	KidClause          string  `orm:"size(255);null"`                   //儿童条款
	AdultBedRate       float32 `orm:"default(0);digits(7);decimals(4)"` //每个成人的加床费
	KidBedRate         float32 `orm:"default(0);digits(7);decimals(4)"` //每个儿童加床费
	BreakfastRate      float32 `orm:"default(0);digits(7);decimals(4)"` //单次早餐价格
	AirportTransfeRate float32 `orm:"default(0);digits(7);decimals(4)"` //单程接机费
	Commission         float32 `orm:"default(0);digits(7);decimals(4)"` //提成比例

}