package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//_, err := gorm.Open("postgres", "goweb:111111@tcp(127.0.0.1:5432)/test?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("postgres", "user=goweb password=111111 dbname=goweb sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connetioon succedssed")
	}
	user := &User{Id: 1, Content: "baba", Author: "xx"}
	db.Create(user)
	//db.Table("abb").Create(user)
	var users User
	//db.Where("Id=?", 1).First(&a)
	db.Find(&users)
	fmt.Println(user)
}

/*
type User struct {
	Id  int   `"boy"`
	Name string   `"gorm:"not_null"`
   }

func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
}
*/
