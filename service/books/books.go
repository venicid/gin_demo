package books

import (
	"fmt"
	"gin_demo/dao"
	"gin_demo/logger"
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
