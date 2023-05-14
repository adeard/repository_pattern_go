package registry

import (
	"gin_web_api/book"

	"gorm.io/gorm"
)

func BookRegistry(db *gorm.DB) book.Service {
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	return bookService
}
