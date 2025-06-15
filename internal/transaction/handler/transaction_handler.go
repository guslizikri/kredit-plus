package handler

import (
	"net/http"
	"strconv"

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

func (h *TransactionHandler) GetTransactionHistory(c *gin.Context) {
	var params dto.GetTransactionHistoryQuery

	// Gunakan c.DefaultQuery untuk memberikan nilai default STRING, lalu konversi ke INT
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)

	if err1 != nil || err2 != nil || page < 1 || limit < 1 || limit > 100 {
		util.SendResponseWithMeta(c, http.StatusBadRequest, nil, nil, "invalid page or limit value")
		return
	}

	params.Page = page
	params.Limit = limit

	role, _ := c.Get("role")
	claimConsumerId, _ := c.Get("consumerId")

	if role.(string) == "consumer" {
		params.ConsumerId = claimConsumerId.(string)
	} else {
		if cid := c.Query("consumer_id"); cid == "" {
			util.SendResponseWithMeta(c, http.StatusBadRequest, nil, nil, "consumer_id required for admin")
			return
		} else {
			params.ConsumerId = cid
		}
	}

	result, total, err := h.usecase.GetTransactionHistory(c, params)
	if err != nil {
		util.SendResponseWithMeta(c, http.StatusInternalServerError, nil, nil, err.Error())
		return
	}

	meta := util.BuildPaginationMeta(params.Page, params.Limit, total)
	util.SendResponseWithMeta(c, http.StatusOK, result, meta, "success get transactions")
}
