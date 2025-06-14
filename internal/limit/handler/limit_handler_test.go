package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/handler"
	"sigmatech-kredit-plus/internal/limit/usecase"
	"sigmatech-kredit-plus/pkg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody    any
		mockUsecaseErr error
		expectedStatus int
	}{
		"successfully create limit": {
			requestBody: dto.SetLimit{
				TenorMonth:  1,
				LimitAmount: 2,
			},
			mockUsecaseErr: nil,
			expectedStatus: http.StatusCreated,
		},
		"usecase returns error": {
			requestBody: dto.SetLimit{
				TenorMonth:  1,
				LimitAmount: 2,
			},
			mockUsecaseErr: errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
		"tenor required": {
			requestBody: dto.SetLimit{
				LimitAmount: 2,
			},
			expectedStatus: http.StatusBadRequest,
		},
		"invalid body": {
			requestBody:    "wrong-format",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUsecase := new(usecase.LimitUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/limit/consumerid", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			mockUsecase.On("SetLimit", mock.Anything, mock.Anything).Return(tc.mockUsecaseErr).Once()

			limitHandler := handler.NewLimitHandler(mockUsecase)
			limitHandler.SetLimit(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
