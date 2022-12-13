package books

import (
	"fmt"
	"gin_demo/controller/common"
	"gin_demo/logger"
	"gin_demo/service/books"
	"github.com/gin-gonic/gin"
)

func GetBookListHandler(c *gin.Context) {
	// 1. 参数处理
	// 2. 业务逻辑
	// 3. 返回响应
	params, err := getListBooksParams(c)
	if err != nil {
		msg := fmt.Sprintf("GetBookListHandler error, error message is %s", err.Error())
		logger.Logger.Error(msg)
		//c.JSON(400, gin.H{"code": 1, "msg": "参数解析错误", "data": nil})
		common.HttpResponse(c, 400, "参数解析错误", nil)
	}

	result, err := books.ListBookRecords(params)
	if err != nil {
		msg := fmt.Sprintf("GetBookListHandler error, error message is %s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 1, "msg": "获取book列表错误", "data": nil})
	}

	c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": result})
}
