package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) Exists(ctx context.Context, consumerID string, tenor int) (bool, error) {
	args := r.Mock.Called(ctx, consumerID, tenor)
	return args.Get(0).(bool), args.Error(1)
}

func (r *RepoMock) CreateLimit(ctx context.Context, limit *model.Limit) error {
	args := r.Mock.Called(ctx, limit)
	return args.Error(0)
}

func (r *RepoMock) UpdateLimit(ctx context.Context, consumerID string, tenor int, limitAmount int) error {
	args := r.Mock.Called(ctx, consumerID, tenor, limitAmount)
	return args.Error(0)
}
