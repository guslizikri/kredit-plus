package usecase_test

import (
	"context"
	"errors"
	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/limit/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLimit(t *testing.T) {
	testCases := map[string]struct {
		mockCreateLimitErr any
		expectedErr        error
	}{
		"successfully create limit": {
			mockCreateLimitErr: nil,
			expectedErr:        nil,
		},
		"error: create limit": {
			mockCreateLimitErr: errors.New("unknown error"),
			expectedErr:        errors.New("unknown error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoLimitMock := new(repository.RepoMock)

			repoLimitMock.On("CreateLimit", mock.Anything, mock.Anything).
				Return(test.mockCreateLimitErr).Once()

			usecase := usecase.NewLimitUsecase(repoLimitMock)

			err := usecase.CreateLimit(context.Background(), &dto.CreateLimit{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
