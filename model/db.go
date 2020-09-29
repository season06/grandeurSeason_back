package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "season1006"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "season"
)

func connect() *gorm.DB {
	CONN := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	var err error
	db, err := gorm.Open("mysql", CONN)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		log.Fatal(err)
	}
	//設置全局表名禁用複數
	db.SingularTable(true)
	db.LogMode(true)

	return db
}
