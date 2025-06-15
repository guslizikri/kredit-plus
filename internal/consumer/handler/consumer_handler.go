package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/usecase"
	"sigmatech-kredit-plus/pkg"
	"sigmatech-kredit-plus/util"
)

type ConsumerHandler struct {
	usecase usecase.ConsumerUsecaseIF
}

func NewConsumerHandler(u usecase.ConsumerUsecaseIF) *ConsumerHandler {
	return &ConsumerHandler{usecase: u}
}

func (h *ConsumerHandler) CreateConsumer(c *gin.Context) {
	var body dto.CreateConsumer
	if err := c.ShouldBind(&body); err != nil {
		util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
		return
	}

	photoKTP, ok := c.Get("photo_ktp")
	if !ok {
		util.SendResponse(c, http.StatusInternalServerError, nil, "photo ktp not found")
		return
	}
	photoSelfie, ok := c.Get("photo_selfie")
	if !ok {
		util.SendResponse(c, http.StatusInternalServerError, nil, "photo selfie not found")
		return
	}
	body.PhotoKTP = photoKTP.(string)
	body.PhotoSelfie = photoSelfie.(string)
	log.Println("2")

	err := pkg.Validate.Struct(&body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			util.SendResponse(c, http.StatusBadRequest, nil, err.Error())
			return
		}
	}

	err = h.usecase.CreateConsumer(c, &body)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusCreated, nil, "success create consumer")
}

func (h *ConsumerHandler) GetConsumerByID(c *gin.Context) {
	id := c.Param("id")
	consumerId, _ := c.Get("consumerId")
	role, _ := c.Get("role")

	if role == "consumer" && consumerId != id {
		util.SendResponse(c, http.StatusForbidden, nil, "error: forbidden")
		return
	}

	consumer, err := h.usecase.GetConsumerDetail(c, id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.SendResponse(c, http.StatusOK, consumer, "success get consumer detail")
}
