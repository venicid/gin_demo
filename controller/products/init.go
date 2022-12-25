package products

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitProducts(e))
	data = append(data, InitTeams(e))
	data = append(data, InitApps(e))
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
	teamApi := e.Group("/api/v1/teams")

	teamApi.GET("")
	teamApi.GET("/:id")
	teamApi.POST("")
	teamApi.PUT("/:id")
	teamApi.DELETE("/:id")

	return teamApi
}

func InitApps(e *gin.Engine) *gin.RouterGroup {
	appApi := e.Group("/api/v1/apps")

	appApi.GET("")
	appApi.GET("/:id")
	appApi.POST("")
	appApi.PUT("/:id")
	appApi.DELETE("/:id")

	return appApi
}
