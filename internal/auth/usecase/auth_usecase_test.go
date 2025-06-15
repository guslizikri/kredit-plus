package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"sigmatech-kredit-plus/internal/auth/dto"
	"sigmatech-kredit-plus/internal/auth/mocks"
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
		inputBody                     *dto.ConsumerLogin
		expectedErr                   error
	}{
		"successfully login": {
			mockGetConsumerByNikReturnErr: nil,
			mockGetConsumerByNikReturnRes: &model.Consumer{
				ID:       "consumer-uuid",
				FullName: "John Doe",
			},
			inputBody: &dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			expectedErr: nil,
		},
		"error: nik doesnt exist": {
			mockGetConsumerByNikReturnErr: sql.ErrNoRows,
			mockGetConsumerByNikReturnRes: nil,
			inputBody: &dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			expectedErr: errors.New("nik doesnt exist"),
		},
		"error: internal server error": {
			mockGetConsumerByNikReturnErr: errors.New("unknown error"),
			mockGetConsumerByNikReturnRes: nil,
			inputBody: &dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			expectedErr: errors.New("unknown error"),
		},
		"error: invalid name": {
			mockGetConsumerByNikReturnErr: nil,
			mockGetConsumerByNikReturnRes: &model.Consumer{
				ID:       "consumer-uuid",
				FullName: "Jane Doe",
			},
			inputBody: &dto.ConsumerLogin{
				NIK:      "1234567890",
				FullName: "John Doe",
			},
			expectedErr: errors.New("name invalid"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			repoAuthMock := new(mocks.RepoMock)

			repoAuthMock.On("GetConsumerByNIK", mock.Anything, test.inputBody.NIK).
				Return(test.mockGetConsumerByNikReturnRes, test.mockGetConsumerByNikReturnErr).Once()

			u := usecase.NewAuthUsecase(repoAuthMock)

			_, err := u.ConsumerLogin(context.Background(), test.inputBody)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			repoAuthMock.AssertExpectations(t)
		})
	}
}

func TestAdminLogin(t *testing.T) {
	u := usecase.NewAuthUsecase(nil)
	body := &dto.AdminLogin{
		Username: "admin",
		Password: "secret",
	}

	token, err := u.AdminLogin(context.Background(), body)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
