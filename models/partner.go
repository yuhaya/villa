package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

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
