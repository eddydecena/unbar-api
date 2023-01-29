package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Pages       int64  `json:"pages"`
}
