package usecase

import (
	"context"
	"sigmatech-kredit-plus/internal/limit/dto"

	"github.com/stretchr/testify/mock"
)

type LimitUsecaseMock struct {
	mock.Mock
}

func (m *LimitUsecaseMock) SetLimit(ctx context.Context, consumerId string, body *dto.SetLimit) error {
	args := m.Called(ctx, body)
	return args.Error(0)
}
