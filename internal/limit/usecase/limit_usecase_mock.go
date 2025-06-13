package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/limit/dto"

	"github.com/stretchr/testify/mock"
)

type LimitUsecaseMock struct {
	mock.Mock
}

func (m *LimitUsecaseMock) CreateLimit(ctx context.Context, body *dto.CreateLimit) error {
	args := m.Called(ctx, body)
	return args.Error(0)
}
