package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateLimit(ctx context.Context, limit *model.Limit) error {
	args := r.Mock.Called(ctx, limit)
	return args.Error(0)
}
