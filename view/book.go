package view

import "gin_demo/model"

type ListBooksRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

type ListBooksResponse struct {
	Count    uint64        `json:"count"`
	Records  []*model.Book `json:"records"`
	Page     int64         `json:"page"`
	PageSize int64         `json:"page_size"`
}
