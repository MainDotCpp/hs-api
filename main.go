package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"hs_short_link/common"
	"hs_short_link/model"
	"hs_short_link/query"
	"hs_short_link/service"
)

func main() {
	// 初始化日志
	common.InitLogger()
	// 初始化数据库
	model.InitDb()
	query.SetDefault(model.DB)
	//genDb()
	r := gin.Default()
	service.InitRouter(r)
	err := r.Run(":9898")
	if err != nil {
		panic(err)
	}
}

func genDb() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(model.DB) // reuse your gorm db
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
