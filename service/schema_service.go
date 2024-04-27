package service

import (
	"github.com/gin-gonic/gin"
	"hs-api/model"
)

func getSchemaById(c *gin.Context) {
	id := c.Param("id")
	var schema model.TSchema
	model.DB.First(&schema, id)
	success(c, schema)
}
func schemeList(c *gin.Context) {
	var schemaList []model.TSchema
	model.DB.Find(&schemaList).Order("id desc")
	success(c, schemaList)
}

func saveSchema(c *gin.Context) {
	var schema model.TSchema
	err := c.BindJSON(&schema)
	if err != nil {
		fail(c, "bind json error")
		return
	}
	if schema.ID == 0 {
		model.DB.Create(&schema)
	} else {
		model.DB.Save(&schema)
	}
	success(c, schema)
}

func deleteSchema(c *gin.Context) {
	var schema model.TSchema
	err := c.BindJSON(&schema)
	if err != nil {
		fail(c, "bind json error")
		return
	}
	model.DB.Delete(&schema)
	c.JSON(200, gin.H{
		"message": "delete schema",
	})
}
