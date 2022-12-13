package books

import (
	"gin_demo/view"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitBooks(e))
	return data
}

func InitBooks(e *gin.Engine) *gin.RouterGroup {
	bookApi := e.Group("/api/v1/books")

	bookApi.GET("", GetBookListHandler)
	bookApi.GET("/:id")
	bookApi.POST("")
	bookApi.PUT("/:id")
	bookApi.DELETE("/:id")

	return bookApi
}

func getListBooksParams(c *gin.Context) (params *view.ListBooksRequest, err error) {
	params = &view.ListBooksRequest{}

	params.Page, params.PageSize = getPagePageSizeParams(c)
	params.SortField, params.SortOrder = getSortFieldDesc(c)

	name := c.Query("name")
	if name != "" {
		params.Name = name
	}

	return params, err
}

func getPagePageSizeParams(c *gin.Context) (pageUint32 int64, pageSizeUint32 int64) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	pageInt32, _ := strconv.Atoi(page)
	pageSizeInt32, _ := strconv.Atoi(pageSize)
	return int64(pageInt32), int64(pageSizeInt32)
}

func getSortFieldDesc(c *gin.Context) (sortField string, sortOrder string) {
	sortField = c.DefaultQuery("sort_field", "id")
	sortOrder = c.DefaultQuery("sort_order", "asc")
	return sortField, sortOrder
}
