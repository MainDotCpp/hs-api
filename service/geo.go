package service

import (
	"github.com/gin-gonic/gin"
	"hs-api/query"
)

func GetGeoList(c *gin.Context) {
	geos, err := query.TGeo.Find()
	if err != nil {
		return
	}
	success(c, geos)
}
