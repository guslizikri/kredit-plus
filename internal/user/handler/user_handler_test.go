package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/user/dto"
	"sigmatech-kredit-plus/internal/user/handler"
	"sigmatech-kredit-plus/internal/user/usecase"
	"sigmatech-kredit-plus/pkg"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody      any
		mockUsecaseErr   error
		setPhotoContext1 bool
		setPhotoContext2 bool
		expectedStatus   int
	}{
		"successfully create user": {
			requestBody: dto.CreateUser{
				NIK:          "1234567890",
				FullName:     "John Doe",
				LegalName:    "John D",
				PlaceOfBirth: "Jakarta",
				DateOfBirth:  time.Now(),
				Salary:       10000000,
			},
			mockUsecaseErr:   nil,
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusCreated,
		},
		"usecase returns error": {
			requestBody: dto.CreateUser{
				NIK:          "999999999",
				FullName:     "Jane Doe",
				LegalName:    "Jane D",
				PlaceOfBirth: "Bandung",
				DateOfBirth:  time.Now(),
				Salary:       9000000,
			},
			mockUsecaseErr:   errors.New("unknown error"),
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusInternalServerError,
		},
		"nik required": {
			requestBody: dto.CreateUser{
				FullName:     "Jane Doe",
				LegalName:    "Jane D",
				PlaceOfBirth: "Bandung",
				DateOfBirth:  time.Now(),
				Salary:       9000000,
			},
			mockUsecaseErr:   errors.New("unknown error"),
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusBadRequest,
		},
		"missing photo context 1": {
			requestBody: dto.CreateUser{
				NIK:          "88888888",
				FullName:     "No Photo",
				LegalName:    "No Photo",
				PlaceOfBirth: "Surabaya",
				DateOfBirth:  time.Now(),
				Salary:       1000000,
			},
			setPhotoContext1: false,
			setPhotoContext2: true,
			expectedStatus:   http.StatusInternalServerError,
		},
		"missing photo context 2": {
			requestBody: dto.CreateUser{
				NIK:          "88888888",
				FullName:     "No Photo",
				LegalName:    "No Photo",
				PlaceOfBirth: "Surabaya",
				DateOfBirth:  time.Now(),
				Salary:       1000000,
			},
			setPhotoContext1: true,
			setPhotoContext2: false,
			expectedStatus:   http.StatusInternalServerError,
		},
		"invalid body": {
			requestBody:    "wrong-format",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUsecase := new(usecase.UserUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			if tc.setPhotoContext1 {
				ctx.Set("photo_ktp", "path/to/ktp.jpg")
			}
			if tc.setPhotoContext2 {
				ctx.Set("photo_selfie", "path/to/selfie.jpg")
			}

			mockUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(tc.mockUsecaseErr).Once()

			userHandler := handler.NewUserHandler(mockUsecase)
			userHandler.CreateUser(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
