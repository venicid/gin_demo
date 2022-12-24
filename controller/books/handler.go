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
		c.JSON(400, gin.H{"code": 400, "msg": "获取book列表错误", "data": nil})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "成功", "data": result})
}

func GetBookDetailHandler(c *gin.Context) {
	bookId, err := common.GetIdParams(c)
	if err != nil {
		msg := fmt.Sprintf("GetBookDetailHandler.GetIdParams.error, error message is %s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	result, err := books.GetBookDetailById(bookId)
	if err != nil {
		msg := fmt.Sprintf("GetBookDetailHandler.error, error message is %s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, fmt.Sprintf("获取book详情错误，id:%v, err:%s", bookId, err.Error()), nil)
		return
	}
	common.HttpResponse(c, 200, "成功", result)

}

func CreateBookHandler(c *gin.Context) {

	params, err := getCreateBookParams(c)
	if err != nil {
		// error：打印校验用户输出错误
		msg := fmt.Sprintf("CreateBookHandler.getCreateBookParams.error.message.is.%s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	err = books.CreateBookRecord(params)
	if err != nil {
		// error：打印错误与用户输入
		msg := fmt.Sprintf("CreateBookHandler.CreateBook.error.message.is.%s.params.%v", err.Error(), params)
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	common.HttpResponse(c, 200, "创建成功", nil)
}

func UpdateBookHandler(c *gin.Context) {
	bookId, err := common.GetIdParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("GetBookDetailHandler.GetIdParams.error, error message is %s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	params, err := getUpdateBookParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("UpdateBookHandler.getUpdateBookParams.error.message.is.%s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	err = books.UpdateBookRecord(bookId, params)
	if err != nil {
		// error：打印错误与用户输入
		msg := fmt.Sprintf("CreateBookHandler.CreateBook.error.message.is.%s.params.%v", err.Error(), params)
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}
	common.HttpResponse(c, 200, "更新成功", nil)
}

func DeleteBookHandler(c *gin.Context) {
	bookId, err := common.GetIdParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("DeleteBookHandler.GetIdParams.error, error message is %s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	err = books.DeleteBookRecord(bookId)
	if err != nil {
		// error：打印错误与用户输入
		msg := fmt.Sprintf("CreateBookHandler.CreateBook.error.message.is.%s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}
	common.HttpResponse(c, 200, "删除成功", nil)
}
