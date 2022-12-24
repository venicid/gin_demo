package books

import (
	"fmt"
	"gin_demo/dao"
	"gin_demo/logger"
	"gin_demo/model"
	"gin_demo/view"
)

func ListBookRecords(req *view.ListBooksRequest) (*view.ListBooksResponse, error) {

	var response = &view.ListBooksResponse{}

	err, result, count := dao.ListBookRecords(req)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.ListBookRecords], result: %v, count:%d, request: %v", result, count, req)
	logger.Logger.Warn(msg)

	response.Count = count
	response.Records = result
	if req.Page == -1 {
		response.Page = 1
		response.PageSize = int64(count)
	} else {
		response.Page = req.Page
		response.PageSize = req.PageSize
	}

	return response, nil
}

func GetBookDetailById(bookId int) (book *model.Book, err error) {

	var result *model.Book
	result, err = dao.GetBookDetailById(bookId)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.GetBookDetailById], bookId: %v,  result: %v", result, bookId)
	logger.Logger.Warn(msg)

	return result, nil
}

func CreateBookRecord(params *view.CreateBookRequest) (err error) {

	book := &model.Book{
		Name: params.Name,
		Desc: params.Desc,
	}
	err = dao.CreateBookRecord(book)
	if err != nil {
		return err
	}

	// warn：打印用户查询到的数据
	msg := fmt.Sprintf("[service.CreateBook], params: %v", params)
	logger.Logger.Warn(msg)

	return nil
}

func UpdateBookRecord(bookId int, params *view.UpdateBookRequest) (err error) {

	book := &model.Book{
		Name: params.Name,
		Desc: params.Desc,
	}
	err = dao.UpdateBookRecord(bookId, book)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("[service.UpdateBookRecord], bookId:%v, params: %v", bookId, params)
	logger.Logger.Warn(msg)

	return nil
}

func DeleteBookRecord(bookId int) (err error) {

	err = dao.DeleteBookRecord(bookId)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("[service.DeleteBookRecord], bookId:%v,", bookId)
	logger.Logger.Warn(msg)

	return nil
}
