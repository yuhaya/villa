package models

type AdminModel struct {
	BaseModel
	Name  string
	Email string
}

//func Main() {
//	session, err := mgo.Dial(beego.AppConfig.String("mgo"))
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()

//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB("test").C("people")
//	err = c.Insert(&Admin{"Armand", "3wmaocomputer@gmail.com"},
//		&Admin{"Yuhaya", "864233002@qq.com"})
//	if err != nil {
//		log.Fatal(err)
//	}

//	result := Admin{}
//	err = c.Find(bson.M{"name": "Armand"}).One(&result)
//	if err != nil {
//		log.Fatal(err)
//	}

//	fmt.Println("Phone:", result.Email)
//}
