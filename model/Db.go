package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	fmt.Printf("init db\n")
	dsn := "host=localhost user=hs_dl password=7ACQpe7yTr73Ej6S dbname=hs_dl port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//err = db.AutoMigrate(&TSchema{}, &TClientInfo{})
	//if err != nil {
	//	return
	//}
	DB = db
	fmt.Print("init db success\n")
	fmt.Print(DB)
}
