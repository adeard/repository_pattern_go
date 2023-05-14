package main

import (
	"fmt"
	"gin_web_api/book"
	"gin_web_api/handler"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "adeardian:@De199429@tcp(127.0.0.1:3306)/gin_web_api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	fmt.Println(bookService.FindAll())

	router := gin.Default()

	v1 := router.Group("v1")
	v1.GET("books", bookHandler.GetBooks)
	v1.POST("books", bookHandler.PostBook)
	v1.GET("books/:ID", bookHandler.GetBook)
	v1.PUT("books/:ID", bookHandler.UpdateBook)
	v1.DELETE("books/:ID", bookHandler.DeleteBook)

	router.Run(":85")
}
