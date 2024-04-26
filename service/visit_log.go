package service

import (
	"github.com/gin-gonic/gin"
	"hs_short_link/query"
	"strconv"
)

type VisitLogService struct{}

func (v VisitLogService) Get(c *gin.Context) {
}

func (v VisitLogService) List(c *gin.Context) {
	no := c.Param("no")
	clients, _ := query.TClientInfo.Where(query.TClientInfo.No.Eq(no)).Find()
	success(c, clients)
}

func (v VisitLogService) Page(c *gin.Context) {
	no := c.Param("no")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))
	offset := (page - 1) * pageSize
	clients, total, _ := query.TClientInfo.Where(query.TClientInfo.No.Eq(no)).FindByPage(offset, pageSize)
	success(c, gin.H{
		"total": total,
		"list":  clients,
	})
}

func (v VisitLogService) Save(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v VisitLogService) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
