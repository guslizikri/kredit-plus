package usecase_test

import (
	"context"
	"errors"
	"sigmatech-kredit-plus/internal/limit/dto"
	"sigmatech-kredit-plus/internal/limit/mocks"
	"sigmatech-kredit-plus/internal/limit/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetLimit(t *testing.T) {
	testCases := map[string]struct {
		mockExistLimitRes  bool
		mockExistLimitErr  any
		mockCreateLimitErr any
		mockUpdateLimitErr any
		expectedErr        error
	}{
		"successfully: create limit": {
			mockExistLimitRes:  false,
			mockExistLimitErr:  nil,
			mockCreateLimitErr: nil,
			expectedErr:        nil,
		},
		"successfully: update limit": {
			mockExistLimitRes:  true,
			mockExistLimitErr:  nil,
			mockUpdateLimitErr: nil,
			expectedErr:        nil,
		},
		"error: check exist": {
			mockExistLimitRes:  false,
			mockExistLimitErr:  errors.New("unknown error"),
			mockCreateLimitErr: nil,
			expectedErr:        errors.New("unknown error"),
		},
		"error: create limit": {
			mockExistLimitRes:  false,
			mockExistLimitErr:  nil,
			mockCreateLimitErr: errors.New("unknown error"),
			expectedErr:        errors.New("unknown error"),
		},
		"error: update limit": {
			mockExistLimitRes:  true,
			mockExistLimitErr:  nil,
			mockUpdateLimitErr: errors.New("unknown error"),
			expectedErr:        errors.New("unknown error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoLimitMock := new(mocks.RepoMock)

			repoLimitMock.On("Exists", mock.Anything, mock.Anything, mock.Anything).
				Return(test.mockExistLimitRes, test.mockExistLimitErr).Once()

			if test.mockExistLimitRes {
				repoLimitMock.On("UpdateLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(test.mockUpdateLimitErr).Once()
			} else {
				repoLimitMock.On("CreateLimit", mock.Anything, mock.Anything).
					Return(test.mockCreateLimitErr).Once()
			}

			usecase := usecase.NewLimitUsecase(repoLimitMock)

			err := usecase.SetLimit(context.Background(), "uuid", &dto.SetLimit{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
