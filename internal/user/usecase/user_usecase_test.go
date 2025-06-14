package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/user/dto"
	"sigmatech-kredit-plus/internal/user/repository"
	"sigmatech-kredit-plus/internal/user/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	testCases := map[string]struct {
		mockGetUserByNikReturnErr any
		mockGetUserByNikReturnRes *model.User
		mockCreateUserErr         any
		expectedErr               error
	}{
		"successfully create user": {
			mockGetUserByNikReturnErr: sql.ErrNoRows,
			mockGetUserByNikReturnRes: &model.User{},
			mockCreateUserErr:         nil,
			expectedErr:               nil,
		},
		"error: create user": {
			mockGetUserByNikReturnErr: sql.ErrNoRows,
			mockGetUserByNikReturnRes: &model.User{},
			mockCreateUserErr:         errors.New("unknown error"),
			expectedErr:               errors.New("unknown error"),
		},
		"error: nik already registered": {
			mockGetUserByNikReturnErr: nil,
			mockGetUserByNikReturnRes: nil,
			expectedErr:               errors.New("nik already registered"),
		},
		"error: get user by nik": {
			mockGetUserByNikReturnErr: errors.New("unknown error"),
			mockGetUserByNikReturnRes: nil,
			expectedErr:               errors.New("unknown error"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoUserMock := new(repository.RepoMock)

			repoUserMock.On("GetUserByNIK", mock.Anything, mock.Anything).
				Return(test.mockGetUserByNikReturnRes, test.mockGetUserByNikReturnErr).Once()

			repoUserMock.On("CreateUser", mock.Anything, mock.Anything).
				Return(test.mockCreateUserErr).Once()

			usecase := usecase.NewUserUsecase(repoUserMock)

			err := usecase.CreateUser(context.Background(), &dto.CreateUser{})

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
