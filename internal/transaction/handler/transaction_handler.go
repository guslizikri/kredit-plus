package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"sigmatech-kredit-plus/internal/transaction/dto"
	"sigmatech-kredit-plus/internal/transaction/usecase"
	"sigmatech-kredit-plus/pkg"
	"sigmatech-kredit-plus/util"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecaseIF
}

func NewTransactionHandler(u usecase.TransactionUsecaseIF) *TransactionHandler {
	return &TransactionHandler{usecase: u}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var body dto.CreateTransaction
	consumerId := c.MustGet("consumerId").(string)
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

	contractNumber, err := h.usecase.CreateTransaction(c, &body, consumerId)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusCreated, contractNumber, "success create transaction")
}
