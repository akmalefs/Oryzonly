package handler

import (
	"net/http"
	"oryzonly/helper"
	"oryzonly/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	response := helper.ResponseJSON("Account has been registered", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}
