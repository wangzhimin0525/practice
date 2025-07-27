package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practice/gorm/example3"
)

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(dst...)
	return db
}

func main() {
	//db := InitDB(&example.Student{})

	//db := InitDB(&example.Account{}, &example.Transaction{})
	//
	//example.Run(db)

	//example2.Run()

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	example3.Run(db)
}
