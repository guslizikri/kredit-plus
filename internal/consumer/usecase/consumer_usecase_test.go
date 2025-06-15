package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/mocks"
	"sigmatech-kredit-plus/internal/consumer/usecase"
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

			repoConsumerMock.On("GetConsumerByNIK", mock.Anything, mock.Anything).
				Return(test.mockGetConsumerByNikReturnRes, test.mockGetConsumerByNikReturnErr).Once()

			repoConsumerMock.On("CreateConsumer", mock.Anything, mock.Anything).
				Return(test.mockCreateConsumerErr).Once()

			usecase := usecase.NewConsumerUsecase(repoConsumerMock)

			err := usecase.CreateConsumer(context.Background(), &dto.CreateConsumer{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
