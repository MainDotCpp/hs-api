package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hs_short_link/model"
)

func Process(c *gin.Context) {
	no := c.Param("no")
	var visitParam model.TClientInfo
	var header = c.Request.Header
	visitParam.IP = c.ClientIP()
	visitParam.No = no
	visitParam.Ua = header.Get("User-Agent")
	visitParam.Lang = header.Get("Accept-Language")
	visitParam.Referer = header.Get("Referer")
	visitParam.VisitTime = c.Request.URL.String()
	fmt.Print(visitParam)

	// 通过no查询数据库
	var schema model.TSchema
	model.DB.Where("no = ?", no).First(&schema)

	result := schema.Check(&visitParam)
	if result {
		c.Redirect(301, schema.DestLink)
		return
	} else {
		c.Status(404)
	}
}
