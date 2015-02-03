package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/validation"
	"time"
	"strings"
)

type Base struct {
}

//管理员表
type Admin struct {
	Id   uint      `orm:"fk;auto"`
	Name string    `orm:"size(20)" valid:"Required"`
	Pwd  string    `orm:"size(40)" valid:"Required"`
	Date time.Time `orm:"auto_now_add" valid:"Required"`
}

func (this *Admin) Valid(v *validation.Validation) {
    if strings.Index(this.Name, "admin") != -1 {
        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
        v.SetError("Name", "名称里不能含有 admin")
    }
}

//当前币种对人民币的汇率
type Currency struct {
	Id           uint      `orm:"fk;auto"`
	Name         string    `orm:"size(12)"`
	ExchangeRate float32   `orm:"digits(7);decimals(4)"`
	UpdateTime   time.Time `orm:"auto_now"`
}

func (c *Currency) TableEngine() string {
	return "MYISAM"
}

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
	Tag        byte   `orm:"default(0)" valid:"Min(0)` //标签 0=>默认
	Status     byte   `orm:"default(1)" valid:"Range(0, 1)` //是否可用 1=>可用 0=>不可用
	Show       byte   `orm:"default(0)" valid:"Range(0, 1)` //是否显示 1=>显示 0=>不显示
}

type DestinationImg struct {
	Id            uint   `orm:"fk;auto"`
	DestinationId uint   `orm:"index"`
	Img           string `orm:size(120)`
	Type          byte   //图片类型
}

type Inquire struct {
	Id              uint      `orm:"fk;auto"`
	ProductId       uint      `orm:"index"`
	DestinationId   uint      `orm:"index"`
	Uid             uint      `orm:"index"`
	StartDate       time.Time `orm:"type(date)"`
	EndDate         time.Time `orm:"type(date)"`
	Status          byte      `orm:"default(0)"`
	UserIp          string    `orm:"size(20);null"`
	AddDate         time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateDate      time.Time `orm:"auto_now;type(datetime)"`
	UserMessage     string    `orm:"size(250);null"` //用户留言
	Memo            string    `orm:"size(512);null"` //处理备注
	Lock            byte      `orm:"default(0)"`     //上锁 0->不上 1-》上
	Source          byte      `orm:"default(0)"`     //订单来源
	Adults          byte      `orm:"default(0)"`
	Kids            byte      `orm:"default(0)"`
	Bedroom         byte      `orm:"default(0)"`
	KidAddBeds      byte      `orm:"default(0)"`
	AdultsAddBeds   byte      `orm:"default(0)"`
	Invoice         byte      `orm:"default(0)"` //是否开发票
	Coupon          byte      `orm:"default(0)"` //是否有优惠劵
	PayDate         time.Time `orm:"null;type(date)"`
	TaxFormula      byte      `orm:"default(0)"` //税费计算公式
	DeprecatedReson string    `orm:"size(250);null"`
}
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

type InquireCurrency struct {
	Id           uint    `orm:"fk;auto"`
	InquireId    uint    `orm:"index"`
	Name         string  `orm:"size(12)"`
	ExchangeRate float32 `orm:"digits(7);decimals(4)"`
}
type InquireExtra struct {
	Id              uint      `orm:"fk;auto"`
	InquireId       uint      `orm:"index"`
	CardId          string    `orm:"size(30);null"`
	FlightNumber    string    `orm:"size(30);null"` //航班号码
	FlightStartTime time.Time `orm:"null;type(datetime)"`
	FlightEndTime   time.Time `orm:"null;type(datetime)"`
}
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
type Partner struct {
	Id                   uint      `orm:"fk;auto" form:"-"`
	Name                 string    `orm:"size(30)"`
	DestinationId        uint      `orm:"default(0)"`
	ManagerContact       string    `orm:"null;size(20)"` //别墅管理联系人
	ManagerTelephone     string    `orm:"null;size(20)"`
	ManagerEmail         string    `orm:"null;size(20)"`
	ReservationContact   string    `orm:"null;size(20)"` //别墅预订联系人
	ReservationTelephone string    `orm:"null;size(20)"`
	ReservationEmail     string    `orm:"null;size(20)"`
	CreatedDate          time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	ContractStartDate    time.Time `orm:"null;auto_now_add;type(date)"` //合作开始时间
	ContractEndDate      time.Time `orm:"null;auto_now_add;type(date)"`
	Commission           float32   `orm:"default(0);digits(7);decimals(4)"` //提成比例
	MembershipGroup      string    `orm:"size(20)"`                         //隶属集团
	State                byte      `orm:"default(0)"`                       //状态 0=》洽谈中 1=》已合作 2=》解除合作 3=》未恢复的 4=》删除的
	Memo                 string    `orm:"size(255);null"`
	Address              string    `orm:"size(255);null"` //联系地址
	Website              string    `orm:"size(100);null"` //网址
	AvailablityLink      string    `orm:"size(100);null"` //别墅可用性查询链接
	AvailablityUsername  string    `orm:"size(30);null"`  //可用性账户
	AvailablityPassword  string    `orm:"size(30);null"`  //可用性密码
}

type Product struct {
	Id                   uint      `orm:"fk;auto"`
	DestinationId        uint      `orm:"default(0)"`
	PartnerId            uint      `orm:"null;default(0)"`
	CurrencyId           byte      `orm:"default(0)"`
	Name                 string    `orm:"size(30)"`
	NameEn               string    `orm:"null;size(50)"`
	Memo                 string    `orm:"null;size(50)"`
	LocationInfo         string    `orm:"null;size(50)"`
	ShowerRoom           byte      `orm:"default(0)"`
	ShortDescription     string    `orm:"null;size(100)"`
	Description          string    `orm:"null;size(500)"`
	Longitude            float32   `orm:"default(0);digits(7);decimals(4)"` //经度
	Latitude             float32   `orm:"default(0);digits(7);decimals(4)"` //纬度
	AddDate              time.Time `orm:"auto_now_add;type(datetime)"`
	ValidFromDate        time.Time `orm:"auto_now_add;type(date);default(0000-00-00)""`
	ValidToDate          time.Time `orm:"auto_now_add;type(date);default(9999-00-00)""`
	State                byte      `orm:"default(0)"`
	CheckinTime          string    `orm:"default();null"`
	CheckOutTime         string    `orm:"default();null"`
	AdultBedRate         float32   `orm:"default(0);digits(7);decimals(4)"` //每个成人的加床费
	KidBedRate           float32   `orm:"default(0);digits(7);decimals(4)"` //每个儿童加床费
	KidAddBedsMax        byte      `orm:"default(0)"`
	AdultsAddBedsMax     byte      `orm:"default(0)"`
	CancellationClause   string    `orm:"size(255);null"`                   //取消预订条款
	KidClause            string    `orm:"size(255);null"`                   //儿童条款
	LocationTaxRadio     float32   `orm:"default(0);digits(7);decimals(4)"` //地方税率
	ServiceTaxRadio      float32   `orm:"default(0);digits(7);decimals(4)"` //服务费税率
	AirportTransfeRate   float32   `orm:"default(0);digits(7);decimals(4)"` //单程接机费
	Commission           float32   `orm:"default(0);digits(7);decimals(4)"` //提成比例
	DepositRatio         float32   `orm:"default(0);digits(7);decimals(4)"` //面积
	TaxFormula           byte      `orm:"default(0)"`                       //税费计算公式
	Website              string    `orm:"size(100);null"`                   //网址
	ReservationContact   string    `orm:"null;size(20)"`                    //别墅预订联系人
	ReservationTelephone string    `orm:"null;size(20)"`
	ReservationEmail     string    `orm:"null;size(20)"`
}
type ProductImg struct {
	Id        uint   `orm:"fk;auto"`
	ProductId uint   `orm:"default(0)"`
	Name      string `orm:"size(30)"`
}
type Rate struct {
	Id        uint      `orm:"fk;auto"`
	ProductId uint      `orm:"default(0)"`
	StartDate time.Time `orm:"type(date)"`
	EndDate   time.Time `orm:"type(date)"`
	Bedroom   byte      `orm:"default(0)"`
	MiniStay  byte      `orm:"default(0)"`
}
type Promotion struct {
	Id uint `orm:"fk;auto"`
}
type Tag struct {
	Id   uint   `orm:"fk;auto"`
	Name string `orm:"size(30)"`
	Type byte   `orm:"default(0)"` //tag类型，1-风格， 2-位置，3-主题，特色 , 4-情景，5-免费服务项目，6-基础设施，7-活动设施，8-综合设施，9-周边
}
type TagProduct struct {
	Id        uint `orm:"fk;auto"`
	ProductId uint `orm:"default(0);index"`
	TagId     uint `orm:"default(0);index"`
}
type User struct {
	Id           uint      `orm:"fk;auto"`
	Name         string    `orm:"size(30)"`
	Password     string    `orm:"size(40)"`
	Email        string    `orm:"size(40);null"`
	Phone        string    `orm:"size(40);null"`
	AddDate      time.Time `orm:"auto_now_add;type(datetime)"`
	UserPassport string    `orm:"size(30);null"`
}

func Regist() {
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName") //数据库名字
	dbPwd := beego.AppConfig.String("dbPwd")
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName) //数据库连接字符串
	orm.RegisterModelWithPrefix("villa_",
		new(Admin),
		new(Currency),
		new(DestinationImg),
		new(InquirePriceCost),
		new(InquireVilla),
		new(InquireCurrency),
		new(InquireExtra),
		new(Destination),
		new(Inquire),
		new(Partner),
		new(Product),
		new(ProductImg),
		new(Promotion),
		new(Rate),
		new(Tag),
		new(TagProduct),
		new(User),
	)
	orm.RegisterDriver("mysql", orm.DR_MySQL)        //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", dbLink) //注册数据库，并设置默认数据库
}

func init() {
	Regist()
	createTable()
}

func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}
