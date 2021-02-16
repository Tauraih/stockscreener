package models 

import ( 
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ..
var DB *gorm.DB
var err error 

// DNS ..
const DNS = "taurai:password@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

// InitialMigration ...
func InitialMigration(){
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to DB")
	}
	DB.AutoMigrate(&User{}, &Company{}, &Prices{})
}
