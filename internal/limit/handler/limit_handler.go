package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/usecase"
	"sigmatech-kredit-plus/pkg"
	"sigmatech-kredit-plus/util"
)

type LimitHandler struct {
	usecase usecase.LimitUsecaseIF
}

func NewLimitHandler(u usecase.LimitUsecaseIF) *LimitHandler {
	return &LimitHandler{usecase: u}
}

func (h *LimitHandler) SetLimit(c *gin.Context) {
	var body dto.SetLimit
	consumerId := c.Param("consumerId")
	if err := c.ShouldBind(&body); err != nil {
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

	err = h.usecase.SetLimit(c, consumerId, &body)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusCreated, nil, "success set limit")
}
