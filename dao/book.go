package dao

import (
	"fmt"
	"gin_demo/dao/mysql"
	"gin_demo/logger"
	"gin_demo/model"
	"gin_demo/view"
)

func ListBookRecords(req *view.ListBooksRequest) (error, []*model.Book, uint64) {

	var queryDb = mysql.GORM
	var records []*model.Book
	var count uint64

	if req.Name != "" {
		queryDb = queryDb.Where("name = ?", req.Name)
	}

	if req.Page > 0 && req.PageSize > 0 {
		queryDb = queryDb.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize)
	}

	err := queryDb.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder)).Find(&records).Count(&count).Error
	if err != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.books.ListBooks sql execute err, err message is %s, req: %v", err, req)
		logger.Logger.Error(msg)
	}
	return err, records, count
}
