package service

import "github.com/gin-gonic/gin"

type RestService interface {
	list(c *gin.Context)
}
