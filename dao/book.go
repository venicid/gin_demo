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

	orderBy := "id desc"
	if req.SortField != "" && req.SortOrder != "" {
		orderBy = fmt.Sprintf("%s %s", req.SortField, req.SortOrder)
	}

	err := queryDb.Order(orderBy).Find(&records).Count(&count).Error
	if err != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.books.ListBooks sql execute err, err message is %s, req: %v", err, req)
		logger.Logger.Error(msg)
	}
	return err, records, count
}

func GetBookDetailById(bookId int) (book *model.Book, err error) {
	book = &model.Book{}
	tx := mysql.GORM.Where("id = ?", bookId).First(&book)
	if tx.Error != nil {
		logger.Logger.Error("【DB.LOG】 dao.books.GetBookDetailById sql execute err, err message is " + tx.Error.Error())
		return nil, tx.Error
	}
	return book, nil
}

func CreateBookRecord(params *model.Book) (err error) {

	tx := mysql.GORM.Create(params)
	if tx.Error != nil {
		// error:打印sql执行的错误
		logger.Logger.Error("【DB.LOG】 dao.books.CreateBook sql execute err, err message is " + tx.Error.Error())
		return tx.Error
	}
	return nil
}

func UpdateBookRecord(bookId int, newBook *model.Book) (err error) {

	tx := mysql.GORM.Model(&model.Book{}).Where("id=?", bookId).Updates(&newBook)
	if tx.Error != nil {
		// error:打印sql执行的错误
		logger.Logger.Error("【DB.LOG】 dao.books.UpdateBookRecord sql execute err, err message is " + tx.Error.Error())
		return tx.Error
	}
	return nil
}

func DeleteBookRecord(bookId int) (err error) {
	tx := mysql.GORM.Where("id=?", bookId).Delete(&model.Book{})
	if tx.Error != nil {
		// error:打印sql执行的错误
		logger.Logger.Error("【DB.LOG】 dao.books.DeleteBookRecord sql execute err, err message is " + tx.Error.Error())
		return tx.Error
	}
	return nil
}
