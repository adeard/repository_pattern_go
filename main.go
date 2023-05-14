package main

import (
	"gin_web_api/config"
	"gin_web_api/handler"
	"gin_web_api/middlewares"
	"gin_web_api/registry"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	bookRegistry := registry.BookRegistry(db)
	bookHandler := handler.NewBookHandler(bookRegistry)
	userRegistry := registry.UserRegistry(db)
	userHandler := handler.NewUserHandler(userRegistry)

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")

	public := v1.Group("public")
	public.POST("/register", userHandler.PostUser)
	public.POST("/login", userHandler.Login)

	v1.Use(middlewares.JwtAuthMiddleware())
	user := v1.Group("users")
	user.GET("/user", userHandler.CurrentUser)
	user.GET("users/:ID", userHandler.GetUser)

	book := v1.Group("books")
	book.GET("", bookHandler.GetBooks)
	book.POST("", bookHandler.PostBook)
	book.GET(":ID", bookHandler.GetBook)
	book.PUT(":ID", bookHandler.UpdateBook)
	book.DELETE(":ID", bookHandler.DeleteBook)

	router.Run(":85")
}
