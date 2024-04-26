package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hs_short_link/common"
)

var DB *gorm.DB

func InitDb() {
	dsn := "host=localhost user=hs_dl password=7ACQpe7yTr73Ej6S dbname=hs_dl port=5432 TimeZone=Asia/Shanghai"
	common.Logger.Infof("初始化数据库: %v", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//err = db.AutoMigrate(&TSchema{}, &TClientInfo{})
	//if err != nil {
	//	return
	//}
	DB = db
	common.Logger.Info("数据库初始化成功")
}
