package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/mocks"
	"sigmatech-kredit-plus/internal/consumer/usecase"
	limit_mocks "sigmatech-kredit-plus/internal/limit/mocks"
	"sigmatech-kredit-plus/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateConsumer(t *testing.T) {
	testCases := map[string]struct {
		mockGetConsumerByNikReturnErr any
		mockGetConsumerByNikReturnRes *model.Consumer
		mockCreateConsumerErr         any
		expectedErr                   error
	}{
		"successfully create consumer": {
			mockGetConsumerByNikReturnErr: sql.ErrNoRows,
			mockGetConsumerByNikReturnRes: &model.Consumer{},
			mockCreateConsumerErr:         nil,
			expectedErr:                   nil,
		},
		"error: create consumer": {
			mockGetConsumerByNikReturnErr: sql.ErrNoRows,
			mockGetConsumerByNikReturnRes: &model.Consumer{},
			mockCreateConsumerErr:         errors.New("unknown error"),
			expectedErr:                   errors.New("unknown error"),
		},
		"error: nik already registered": {
			mockGetConsumerByNikReturnErr: nil,
			mockGetConsumerByNikReturnRes: nil,
			expectedErr:                   errors.New("nik already registered"),
		},
		"error: get consumer by nik": {
			mockGetConsumerByNikReturnErr: errors.New("unknown error"),
			mockGetConsumerByNikReturnRes: nil,
			expectedErr:                   errors.New("unknown error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoConsumerMock := new(mocks.RepoMock)
			repoLimitMock := new(limit_mocks.RepoMock)

			repoConsumerMock.On("GetConsumerByNIK", mock.Anything, mock.Anything).
				Return(test.mockGetConsumerByNikReturnRes, test.mockGetConsumerByNikReturnErr).Once()

			repoConsumerMock.On("CreateConsumer", mock.Anything, mock.Anything).
				Return(test.mockCreateConsumerErr).Once()

			usecase := usecase.NewConsumerUsecase(repoConsumerMock, repoLimitMock)

			err := usecase.CreateConsumer(context.Background(), &dto.CreateConsumer{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
func TestGetConsumerDetail(t *testing.T) {
	testCases := map[string]struct {
		mockGetConsumerByIDReturnErr error
		mockGetConsumerByIDReturnRes *dto.GetConsumerDetailResponse
		mockGetLimitReturnRes        []*model.Limit
		mockGetLimitReturnErr        error
		expectedErr                  error
	}{
		"successfully get consumer detail": {
			mockGetConsumerByIDReturnErr: nil,
			mockGetConsumerByIDReturnRes: &dto.GetConsumerDetailResponse{
				ID:       "consumer-uuid",
				FullName: "John Doe",
			},
			mockGetLimitReturnRes: []*model.Limit{
				{
					ID:          "limit-uuid",
					ConsumerID:  "consumer-uuid",
					TenorMonth:  12,
					LimitAmount: 1000000,
					UsedAmount:  200000,
				},
			},
			expectedErr: nil,
		},
		"error: consumer not found": {
			mockGetConsumerByIDReturnErr: sql.ErrNoRows,
			mockGetConsumerByIDReturnRes: nil,
			expectedErr:                  errors.New("consumer not found"),
		},
		"error: get consumer by id": {
			mockGetConsumerByIDReturnErr: errors.New("unknown error"),
			mockGetConsumerByIDReturnRes: nil,
			expectedErr:                  errors.New("unknown error"),
		},
		"error: get limit by consumer": {
			mockGetConsumerByIDReturnErr: nil,
			mockGetConsumerByIDReturnRes: &dto.GetConsumerDetailResponse{
				ID:       "consumer-uuid",
				FullName: "John Doe",
			},
			mockGetLimitReturnErr: errors.New("limit db error"),
			expectedErr:           errors.New("limit db error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoConsumerMock := new(mocks.RepoMock)
			repoLimitMock := new(limit_mocks.RepoMock)

			repoConsumerMock.On("GetConsumerByID", mock.Anything, mock.Anything).
				Return(test.mockGetConsumerByIDReturnRes, test.mockGetConsumerByIDReturnErr).Once()

			if test.mockGetConsumerByIDReturnErr == nil && test.mockGetConsumerByIDReturnRes != nil {
				repoLimitMock.On("GetLimitByConsumerID", mock.Anything, test.mockGetConsumerByIDReturnRes.ID).
					Return(test.mockGetLimitReturnRes, test.mockGetLimitReturnErr).Once()
			}

			u := usecase.NewConsumerUsecase(repoConsumerMock, repoLimitMock)

			_, err := u.GetConsumerDetail(context.Background(), "consumer-uuid")

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			repoConsumerMock.AssertExpectations(t)
			repoLimitMock.AssertExpectations(t)
		})
	}
}
