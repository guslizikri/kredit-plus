package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/transaction/dto"
	"sigmatech-kredit-plus/internal/transaction/handler"
	"sigmatech-kredit-plus/internal/transaction/mocks"
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
			mockUsecase := new(mocks.TransactionUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/transaction/consumerid", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("consumerId", "consumer uuid")

			mockUsecase.On("CreateTransaction", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.mockUsecaseRes, tc.mockUsecaseErr).Once()

			transactionHandler := handler.NewTransactionHandler(mockUsecase)
			transactionHandler.CreateTransaction(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestGetTransactionHistory(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := map[string]struct {
		role               string
		consumerID         string
		query              string
		mockResult         []*dto.GetTransactionHistoryResponse
		mockTotal          int
		mockErr            error
		expectedStatusCode int
		expectedMsg        string
	}{
		"success consumer": {
			role:       "consumer",
			consumerID: "consumer-uuid",
			query:      "",
			mockResult: []*dto.GetTransactionHistoryResponse{
				{ContractNumber: "TRX-123"},
			},
			mockTotal:          1,
			mockErr:            nil,
			expectedStatusCode: http.StatusOK,
			expectedMsg:        "success get transactions",
		},
		"success admin": {
			role:  "admin",
			query: "?page=1&limit=10&consumer_id=admin-consumer-id",
			mockResult: []*dto.GetTransactionHistoryResponse{
				{ContractNumber: "TRX-ADMIN"},
			},
			mockTotal:          1,
			mockErr:            nil,
			expectedStatusCode: http.StatusOK,
			expectedMsg:        "success get transactions",
		},
		"error: invalid page param": {
			role:               "consumer",
			consumerID:         "consumer-uuid",
			query:              "?page=abc&limit=10",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        "invalid page or limit value",
		},
		"error: consumer_id required for admin": {
			role:               "admin",
			query:              "",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        "consumer_id required for admin",
		},
		"error: usecase failure": {
			role:               "consumer",
			consumerID:         "consumer-uuid",
			query:              "?page=1&limit=10",
			mockErr:            errors.New("db error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        "db error",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUC := new(mocks.TransactionUsecaseMock)

			if tc.expectedStatusCode == http.StatusOK || tc.mockErr != nil {
				mockUC.On("GetTransactionHistory", mock.Anything, mock.Anything).
					Return(tc.mockResult, tc.mockTotal, tc.mockErr).Once()
			}

			h := handler.NewTransactionHandler(mockUC)

			r := gin.New()
			r.GET("/transactions", func(c *gin.Context) {
				c.Set("role", tc.role)
				c.Set("consumerId", tc.consumerID)
				h.GetTransactionHistory(c)
			})

			req := httptest.NewRequest(http.MethodGet, "/transactions"+tc.query, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedMsg)

			mockUC.AssertExpectations(t)
		})
	}
}
