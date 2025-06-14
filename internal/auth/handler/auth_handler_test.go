package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/auth/dto"
	"sigmatech-kredit-plus/internal/auth/handler"
	"sigmatech-kredit-plus/internal/auth/usecase"
	"sigmatech-kredit-plus/pkg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConsumerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody      any
		mockUsecaseRes   string
		mockUsecaseErr   error
		setPhotoContext1 bool
		setPhotoContext2 bool
		expectedStatus   int
	}{
		"successfully login": {
			requestBody: dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			mockUsecaseRes: "token",
			mockUsecaseErr: nil,
			expectedStatus: http.StatusOK,
		},
		"error: usecase error": {
			requestBody: dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			mockUsecaseRes: "token",
			mockUsecaseErr: errors.New("internal error"),
			expectedStatus: http.StatusInternalServerError,
		},
		"error: nik required": {
			requestBody: dto.ConsumerLogin{
				FullName: "John Doe",
			},
			expectedStatus: http.StatusBadRequest,
		},
		"error: body parser": {
			requestBody:    "wrong-format",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUsecase := new(usecase.AuthUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			mockUsecase.On("ConsumerLogin", mock.Anything, mock.Anything).Return(tc.mockUsecaseRes, tc.mockUsecaseErr).Once()

			authHandler := handler.NewAuthHandler(mockUsecase)
			authHandler.ConsumerLogin(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAdminLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody      any
		mockUsecaseRes   string
		mockUsecaseErr   error
		setPhotoContext1 bool
		setPhotoContext2 bool
		expectedStatus   int
	}{
		"successfully login": {
			requestBody: dto.AdminLogin{
				Username: "1234567890",
				Password: "John Doe",
			},
			mockUsecaseRes: "token",
			mockUsecaseErr: nil,
			expectedStatus: http.StatusOK,
		},
		"error: usecase error": {
			requestBody: dto.AdminLogin{
				Username: "1234567890",
				Password: "John Doe",
			},
			mockUsecaseRes: "token",
			mockUsecaseErr: errors.New("internal error"),
			expectedStatus: http.StatusInternalServerError,
		},
		"error: username required": {
			requestBody: dto.AdminLogin{
				Password: "John Doe",
			},
			expectedStatus: http.StatusBadRequest,
		},
		"error: body parser": {
			requestBody:    "wrong-format",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUsecase := new(usecase.AuthUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			mockUsecase.On("AdminLogin", mock.Anything, mock.Anything).Return(tc.mockUsecaseRes, tc.mockUsecaseErr).Once()

			authHandler := handler.NewAuthHandler(mockUsecase)
			authHandler.AdminLogin(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
