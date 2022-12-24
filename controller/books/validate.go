package books

import (
	"errors"
	"gin_demo/controller/common"
	"gin_demo/view"
	"github.com/gin-gonic/gin"
)

/**
book相关的
*/
func getListBooksParams(c *gin.Context) (params *view.ListBooksRequest, err error) {
	params = &view.ListBooksRequest{}

	params.Page, params.PageSize = common.GetPagePageSizeParams(c)
	params.SortField, params.SortOrder = common.GetSortFieldDesc(c)

	name := c.Query("name")
	if name != "" {
		params.Name = name
	}
	return params, err
}

func getCreateBookParams(c *gin.Context) (params *view.CreateBookRequest, err error) {
	params = &view.CreateBookRequest{}

	if err := c.ShouldBindJSON(&params); err != nil {
		return nil, errors.New("bind绑定参数失败")
	}

	if params.Name == "" {
		return nil, errors.New("name不能为空")
	}

	return params, err
}

func getUpdateBookParams(c *gin.Context) (params *view.UpdateBookRequest, err error) {
	params = &view.UpdateBookRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		return nil, errors.New("bind绑定参数失败")
	}

	if params.Name == "" {
		return nil, errors.New("name不能为空")
	}
	return params, err
}
