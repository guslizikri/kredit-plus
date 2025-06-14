package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"sigmatech-kredit-plus/internal/auth/dto"
	"sigmatech-kredit-plus/internal/auth/usecase"
	"sigmatech-kredit-plus/pkg"
	"sigmatech-kredit-plus/util"
)

type AuthHandler struct {
	usecase usecase.AuthUsecaseIF
}

func NewAuthHandler(u usecase.AuthUsecaseIF) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) ConsumerLogin(c *gin.Context) {
	var body dto.ConsumerLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	err := pkg.Validate.Struct(&body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
			return
		}
	}

	token, err := h.usecase.ConsumerLogin(c, &body)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusOK, token, "login success")
}

func (h *AuthHandler) AdminLogin(c *gin.Context) {
	var body dto.AdminLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	err := pkg.Validate.Struct(&body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
			return
		}
	}

	token, err := h.usecase.AdminLogin(c, &body)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusOK, token, "login success")
}
