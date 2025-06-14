package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/auth/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUsecaseMock struct {
	mock.Mock
}

func (m *AuthUsecaseMock) ConsumerLogin(ctx context.Context, body *dto.ConsumerLogin) (token string, err error) {
	args := m.Called(ctx, body)
	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUsecaseMock) AdminLogin(ctx context.Context, body *dto.AdminLogin) (token string, err error) {
	args := m.Called(ctx, body)
	return args.Get(0).(string), args.Error(1)
}
