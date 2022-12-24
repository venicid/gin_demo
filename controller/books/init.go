package books

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitBooks(e))
	return data
}

func InitBooks(e *gin.Engine) *gin.RouterGroup {
	bookApi := e.Group("/api/v1/books")

	bookApi.GET("", GetBookListHandler)
	bookApi.GET("/:id", GetBookDetailHandler)
	bookApi.POST("", CreateBookHandler)
	bookApi.PUT("/:id", UpdateBookHandler)
	bookApi.DELETE("/:id", DeleteBookHandler)

	return bookApi
}
