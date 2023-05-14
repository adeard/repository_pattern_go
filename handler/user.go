package handler

import (
	"fmt"
	"gin_web_api/user"
	"gin_web_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *gin.Context) {
	var loginInput user.LoginRequest

	c.ShouldBind(&loginInput)

	token, err := h.userService.Login(loginInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	u, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

func (h *userHandler) PostUser(c *gin.Context) {
	var userInput user.RegisterRequest

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})
	}

	user, err := h.userService.Create(userInput)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.FindByID(int(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
