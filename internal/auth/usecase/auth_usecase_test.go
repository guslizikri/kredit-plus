package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/auth/dto"
	"sigmatech-kredit-plus/internal/auth/repository"
	"sigmatech-kredit-plus/internal/auth/usecase"
	"sigmatech-kredit-plus/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConsumerLogin(t *testing.T) {
	testCases := map[string]struct {
		mockGetConsumerByNikReturnErr any
		mockGetConsumerByNikReturnRes *model.Consumer
		mockConsumerLoginErr          any
		expectedErr                   error
	}{
		"successfully login": {
			mockGetConsumerByNikReturnErr: nil,
			mockGetConsumerByNikReturnRes: &model.Consumer{},
			mockConsumerLoginErr:          nil,
			expectedErr:                   nil,
		},
		"error: nik doesnt exist": {
			mockGetConsumerByNikReturnErr: sql.ErrNoRows,
			mockGetConsumerByNikReturnRes: nil,
			expectedErr:                   errors.New("nik doesnt exist"),
		},
		"error: internal server error": {
			mockGetConsumerByNikReturnErr: errors.New("unknown error"),
			mockGetConsumerByNikReturnRes: nil,
			expectedErr:                   errors.New("unknown error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoAuthMock := new(repository.RepoMock)

			repoAuthMock.On("GetConsumerByNIK", mock.Anything, mock.Anything).
				Return(test.mockGetConsumerByNikReturnRes, test.mockGetConsumerByNikReturnErr).Once()

			usecase := usecase.NewAuthUsecase(repoAuthMock)

			_, err := usecase.ConsumerLogin(context.Background(), &dto.ConsumerLogin{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
