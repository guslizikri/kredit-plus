package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sigmatech-kredit-plus/internal/user/dto"
	"sigmatech-kredit-plus/internal/user/usecase"
	"sigmatech-kredit-plus/util"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	util.SendResponse(c, http.StatusCreated, nil, "success create user")
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.SendResponse(c, http.StatusOK, user, "success get user detail")
}
