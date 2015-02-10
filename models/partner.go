package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

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
	ContractStartDate    time.Time `orm:"null;auto_now_add;type(date)" form:"ContractStartDate,2006-01-02"` //合作开始时间
	ContractEndDate      time.Time `orm:"null;auto_now_add;type(date)" form:"ContractEndDate,2006-01-02"`
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

type PartnerModel struct {
}

func (this *PartnerModel) ListByStatus(status byte) ([]*Partner, int64, error) {

	o := orm.NewOrm()
	partner := new(Partner)
	qs := o.QueryTable(partner)
	var partners []*Partner
	num, err := qs.Filter("State", status).OrderBy("CreatedDate").All(&partners)
	return partners, num, err

}
