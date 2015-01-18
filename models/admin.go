package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type AdminModel struct {
}

func (this *AdminModel) Login(v_name string, v_pwd string) (bool, error, int) {
	o := orm.NewOrm()
	t_admin := Admin{Name: v_name, Pwd: v_pwd}
	err := o.Read(&t_admin, "Name", "Pwd")
	if err != nil {
		return false, err, 0
	} else {
		fmt.Println(t_admin.Id, t_admin.Name)
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
