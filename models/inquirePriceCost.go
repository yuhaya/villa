package models

type InquirePriceCost struct {
	Id                  uint    `orm:"fk;auto"`
	InquireId           uint    `orm:"index"`
	InvoicePrice        float32 `orm:"default(0);digits(7);decimals(4)"` //发票价格
	CouponPrice         float32 `orm:"default(0);digits(7);decimals(4)"` //优惠劵
	FlightsIn           float32 `orm:"default(0);digits(7);decimals(4)"` //机票收入
	FlightsOut          float32 `orm:"default(0);digits(7);decimals(4)"` //机票支出
	InsuranceIn         float32 `orm:"default(0);digits(7);decimals(4)"` //保险收入
	InsuranceOut        float32 `orm:"default(0);digits(7);decimals(4)"` //保险支出
	PoundagePay         float32 `orm:"default(0);digits(7);decimals(4)"` //支付手续费
	PoundageRemit       float32 `orm:"default(0);digits(7);decimals(4)"` //汇款手续费
	AdultsBedRate       float32 `orm:"default(0);digits(7);decimals(4)"` //成人加床费
	KidsBedRate         float32 `orm:"default(0);digits(7);decimals(4)"` //儿童加床费
	RoomsRateIn         float32 `orm:"default(0);digits(7);decimals(4)"` //房费收入
	RoomsRateOut        float32 `orm:"default(0);digits(7);decimals(4)"` //房费支出
	AirportTransfesRate float32 `orm:"default(0);digits(7);decimals(4)"` //接机费
	LocationRate        float32 `orm:"default(0);digits(7);decimals(4)"` //地方税
	ServiceRate         float32 `orm:"default(0);digits(7);decimals(4)"` //服务费
	BreakfastsRate      float32 `orm:"default(0);digits(7);decimals(4)"` //早餐总价格
}