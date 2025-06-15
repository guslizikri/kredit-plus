package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/common"
	limit_mocks "sigmatech-kredit-plus/internal/limit/mocks"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/transaction/dto"
	"sigmatech-kredit-plus/internal/transaction/mocks"
	"sigmatech-kredit-plus/internal/transaction/usecase"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction_WithTransactionManagerMock(t *testing.T) {
	testCases := map[string]struct {
		mockBeginErr             error
		mockGetLimit             *model.Limit
		mockGetLimitErr          error
		mockUpdateUsedErr        error
		mockCreateTransactionErr error
		mockCommitErr            error
		expectedErr              error
	}{
		"successfully": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 1000000,
				UsedAmount:  0,
			},
		},
		"error: Limit tenor not found for this consumer": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 1000000,
				UsedAmount:  0,
			},
			mockGetLimitErr: sql.ErrNoRows,
			expectedErr:     errors.New("Limit tenor not found for this consumer"),
		},
		"error: begin transaction": {
			mockBeginErr: errors.New("begin error"),
			expectedErr:  errors.New("begin error"),
		},
		"error: get limit": {
			mockGetLimitErr: errors.New("limit not found"),
			expectedErr:     errors.New("limit not found"),
		},
		"error: insufficient limit": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 100000,
				UsedAmount:  0,
			},
			expectedErr: errors.New("insufficient limit"),
		},
		"error: update used_amount": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 1000000,
				UsedAmount:  0,
			},
			mockUpdateUsedErr: errors.New("failed update"),
			expectedErr:       errors.New("failed update"),
		},
		"error: create transaction": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 1000000,
				UsedAmount:  0,
			},
			mockCreateTransactionErr: errors.New("insert failed"),
			expectedErr:              errors.New("insert failed"),
		},
		"error: commit failed": {
			mockGetLimit: &model.Limit{
				ID:          "limit-uuid",
				ConsumerID:  "consumer-uuid",
				TenorMonth:  3,
				LimitAmount: 1000000,
				UsedAmount:  0,
			},
			mockCommitErr: errors.New("commit failed"),
			expectedErr:   errors.New("commit failed"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			trxManager := new(common.TransactionManagerMock)
			limitRepo := new(limit_mocks.RepoMock)
			trxRepo := new(mocks.RepoMock)
			txMock := &sqlx.Tx{}

			if test.mockBeginErr != nil {
				trxManager.On("Begin", mock.Anything).Return((*sqlx.Tx)(nil), test.mockBeginErr).Once()
			} else {
				trxManager.On("Begin", mock.Anything).Return(txMock, nil).Once()
				trxManager.On("Rollback", mock.Anything, txMock).Return(nil).Maybe()
			}

			if test.mockBeginErr == nil {
				limitRepo.On("GetLimitWithLock", mock.Anything, txMock, mock.Anything, mock.Anything).
					Return(test.mockGetLimit, test.mockGetLimitErr).Once()
			}

			if test.mockBeginErr == nil && test.mockGetLimitErr == nil && test.mockGetLimit != nil && (test.expectedErr == nil || test.expectedErr.Error() != "insufficient limit") {
				limitRepo.On("UpdateUsedAmountWithTx", mock.Anything, txMock, test.mockGetLimit.ID, mock.Anything).
					Return(test.mockUpdateUsedErr).Maybe()
			}

			if test.mockBeginErr == nil && test.mockGetLimitErr == nil && test.mockUpdateUsedErr == nil && test.mockGetLimit != nil && (test.mockCreateTransactionErr != nil || test.mockCommitErr != nil || test.expectedErr == nil) {
				trxRepo.On("CreateTransactionWithTx", mock.Anything, txMock, mock.Anything).
					Return(test.mockCreateTransactionErr).Once()
			}

			if test.mockBeginErr == nil && test.mockGetLimitErr == nil && test.mockUpdateUsedErr == nil && test.mockCreateTransactionErr == nil && test.mockGetLimit != nil {
				trxManager.On("Commit", mock.Anything, txMock).Return(test.mockCommitErr).Maybe()
			}

			usecase := usecase.NewTransactionUsecase(trxManager, limitRepo, trxRepo)

			body := &dto.CreateTransaction{
				TenorMonth:  3,
				OTRPrice:    500000,
				AdminFee:    10000,
				Installment: 100000,
				Interest:    15000,
				AssetName:   "Test Product",
			}

			_, err := usecase.CreateTransaction(context.Background(), body, "consumer-uuid")

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			trxManager.AssertExpectations(t)
			limitRepo.AssertExpectations(t)
			trxRepo.AssertExpectations(t)
		})
	}
}

func TestGetTransactionHistory(t *testing.T) {
	testCases := map[string]struct {
		mockFetchResult []*dto.GetTransactionHistoryResponse
		mockFetchErr    error
		mockCountResult int
		mockCountErr    error
		expectedErr     error
	}{
		"successfully": {
			mockFetchResult: []*dto.GetTransactionHistoryResponse{
				{
					ContractNumber: "TRX-123",
					TenorMonth:     12,
					OTRPrice:       1000000,
					Installment:    100000,
					Interest:       50000,
					AssetName:      "Motor",
					CreatedAt:      time.Now(),
				},
			},
			mockFetchErr:    nil,
			mockCountResult: 1,
			expectedErr:     nil,
		},
		"error: fetch transactions": {
			mockFetchErr: errors.New("db error"),
			expectedErr:  errors.New("db error"),
		},
		"error: count transactions": {
			mockFetchResult: []*dto.GetTransactionHistoryResponse{},
			mockCountErr:    errors.New("count error"),
			expectedErr:     errors.New("count error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			trxManager := new(common.TransactionManagerMock)
			limitRepo := new(limit_mocks.RepoMock)
			trxRepo := new(mocks.RepoMock)
			// txMock := &sqlx.Tx{}

			trxRepo.On("FetchTransactionByConsumer", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(test.mockFetchResult, test.mockFetchErr).Once()

			if test.mockFetchErr == nil {
				trxRepo.On("CountTransactionByConsumer", mock.Anything, mock.Anything).
					Return(test.mockCountResult, test.mockCountErr).Once()
			}

			u := usecase.NewTransactionUsecase(trxManager, limitRepo, trxRepo)

			params := dto.GetTransactionHistoryQuery{
				Page:       1,
				Limit:      10,
				ConsumerId: "consumer-uuid",
			}

			_, _, err := u.GetTransactionHistory(context.Background(), params)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			trxManager.AssertExpectations(t)
			limitRepo.AssertExpectations(t)
			trxRepo.AssertExpectations(t)
		})
	}
}
