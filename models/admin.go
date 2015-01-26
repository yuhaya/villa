package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type AdminModel struct {
}

func (this *AdminModel) Login(v_name string, v_pwd string) (bool, error, uint) {
	o := orm.NewOrm()
	v_pwd_encrypt_byte := md5.Sum([]byte(v_pwd))
	v_pwd_encrypt := string(v_pwd_encrypt_byte[:])
	t_admin := Admin{Name: v_name, Pwd: v_pwd_encrypt}
	fmt.Printf("username:%v  password:%v\n", v_name, v_pwd_encrypt_byte)
	fmt.Printf("username:%v  password:%v\n", v_name, v_pwd_encrypt_byte[:])
	fmt.Printf("username:%v  password:%v\n", v_name, v_pwd_encrypt)
	err := o.Read(&t_admin, "Name", "Pwd")
	if err != nil {
		return false, err, 0
	} else {
		return true, nil, t_admin.Id
	}
}
func main() {
	//	o := orm.NewOrm()
	//
	//	user := Admin{Name: "slene"}
	//
	//	// insert
	//	id, err := o.Insert(&user)
	//	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//	// update
	//	user.Name = "astaxie"
	//	num, err := o.Update(&user)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//	// read one
	//	u := User{Id: user.Id}
	//	err = o.Read(&u)
	//	fmt.Printf("ERR: %v\n", err)
	//
	//	// delete
	//	num, err = o.Delete(&u)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
