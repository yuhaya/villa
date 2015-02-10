package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/validation"
	"time"
	"strings"
)

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

type AdminModel struct {
}

func (this *AdminModel) Login(v_name string, v_pwd string) (bool, error, uint) {
	o := orm.NewOrm()
	v_pwd_encrypt_byte := md5.Sum([]byte(v_pwd))
	v_pwd_encrypt := hex.EncodeToString(v_pwd_encrypt_byte[:])
	t_admin := Admin{Name: v_name, Pwd: v_pwd_encrypt}
	fmt.Printf("username:%s  password:%s\n", v_name, v_pwd_encrypt)
	err := o.Read(&t_admin, "Name", "Pwd")
	if err != nil {
		return false, err, 0
	} else {
		return true, nil, t_admin.Id
	}
}

func (this *AdminModel) UserInfo(admin_id uint) (Admin, error) {

	o := orm.NewOrm()
	u_admin := Admin{Id: admin_id}
	err := o.Read(&u_admin)
	return u_admin, err
}