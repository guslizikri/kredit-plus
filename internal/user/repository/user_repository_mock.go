package repository

import (
	"context"
	"sigmatech-kredit-plus/internal/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateUser(ctx context.Context, user *model.User) error {
	args := r.Mock.Called(ctx, user)
	return args.Error(0)
}

func (r *RepoMock) GetUserByNIK(ctx context.Context, nik string) (*model.User, error) {
	args := r.Mock.Called(ctx, nik)
	return args.Get(0).(*model.User), args.Error(1)
}
