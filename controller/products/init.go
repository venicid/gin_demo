package products

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitProducts(e))
	data = append(data, InitTeams(e))
	return data
}

func InitProducts(e *gin.Engine) *gin.RouterGroup {
	productApi := e.Group("/api/v1/products")
	productApi.GET("")
	productApi.GET("/:id")
	productApi.POST("")
	productApi.PUT("/:id")
	productApi.DELETE("/:id")
	return productApi
}

func InitTeams(e *gin.Engine) *gin.RouterGroup {
	productApi := e.Group("/api/v1/teams")
	productApi.GET("")
	return productApi
}
