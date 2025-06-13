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
	var body dto.CreateUser
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

	err := h.usecase.CreateUser(c, &body)
	if err != nil {
		e := util.ToHttpError(err)
		util.SendResponse(c, e.Code, nil, e.Error())
		return
	}

	util.SendResponse(c, http.StatusCreated, nil, "success create user")
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUserByNIK(c, id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	util.SendResponse(c, http.StatusOK, user, "success get user detail")
}
