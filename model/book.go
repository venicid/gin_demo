package model

type Book struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" binding:"required" json:"name"`
	Desc string `json:"desc"`
}

func (receiver Book) TableName() string {
	return "book"
}
