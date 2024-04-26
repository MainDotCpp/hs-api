package service

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	r.GET(":no", Process)
	schemaRouter := r.Group("/schema")
	{
		schemaRouter.GET("/:id", getSchemaById)
		schemaRouter.GET("/list", schemeList)
		schemaRouter.POST("/save", saveSchema)
		schemaRouter.POST("/delete", deleteSchema)
	}
	geoRouter := r.Group("/geo")
	{
		geoRouter.GET("/list", GetGeoList)
	}

	visitLogRouter := r.Group("/visit_log")
	{
		visitLogRouter.GET("/:no", VisitLogService{}.List)
		visitLogRouter.GET("/:no/page", VisitLogService{}.Page)
	}
}

func success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
		"code":    200,
	})
}

func fail(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"message": message,
		"code":    500,
	})
}
