package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/user/dto"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) CreateUser(ctx context.Context, body *dto.CreateUser) error {
	args := m.Called(ctx, body)
	return args.Error(0)
}

func (m *UserUsecaseMock) GetUserByNIK(ctx context.Context, nik string) (*model.User, error) {
	args := m.Called(ctx, nik)
	return args.Get(0).(*model.User), args.Error(1)
}
