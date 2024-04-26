package service

import "github.com/gin-gonic/gin"

type RestService interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	Page(c *gin.Context)
	Save(c *gin.Context)
	Delete(c *gin.Context)
}
