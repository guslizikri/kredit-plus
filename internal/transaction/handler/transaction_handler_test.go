package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/transaction/dto"
	"sigmatech-kredit-plus/internal/transaction/handler"
	"sigmatech-kredit-plus/internal/transaction/usecase"
	"sigmatech-kredit-plus/pkg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody    any
		mockUsecaseRes string
		mockUsecaseErr error
		expectedStatus int
	}{
		"successfully create transaction": {
			requestBody: dto.CreateTransaction{
				TenorMonth:    2,
				OTRPrice:      300000,
				AdminFee:      5000,
				Installment:   15000,
				Interest:      5000,
				AssetName:     "Motor",
				SourceChannel: "Dealer",
			},
			mockUsecaseErr: nil,
			mockUsecaseRes: "asd",
			expectedStatus: http.StatusCreated,
		},
		"usecase returns error": {
			requestBody: dto.CreateTransaction{
				TenorMonth:    2,
				OTRPrice:      300000,
				AdminFee:      5000,
				Installment:   15000,
				Interest:      5000,
				AssetName:     "Motor",
				SourceChannel: "Dealer",
			},
			mockUsecaseErr: errors.New("unknown error"),
			mockUsecaseRes: "",
			expectedStatus: http.StatusInternalServerError,
		},
		"tenor required": {
			requestBody: dto.CreateTransaction{
				OTRPrice:      300000,
				AdminFee:      5000,
				Installment:   15000,
				Interest:      5000,
				AssetName:     "Motor",
				SourceChannel: "Dealer",
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
			mockUsecase := new(usecase.TransactionUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/transaction/consumerid", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			mockUsecase.On("CreateTransaction", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.mockUsecaseRes, tc.mockUsecaseErr).Once()

			transactionHandler := handler.NewTransactionHandler(mockUsecase)
			transactionHandler.CreateTransaction(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
