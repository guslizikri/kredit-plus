package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sigmatech-kredit-plus/internal/common"
	limit_repository "sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/model"
	"sigmatech-kredit-plus/internal/transaction/dto"
	"sigmatech-kredit-plus/internal/transaction/repository"
	"sigmatech-kredit-plus/util"
	"time"

	"github.com/google/uuid"
)

type TransactionUsecaseIF interface {
	CreateTransaction(ctx context.Context, body *dto.CreateTransaction, consumerId string) (string, error)
	GetTransactionHistory(ctx context.Context, params dto.GetTransactionHistoryQuery) ([]*dto.GetTransactionHistoryResponse, int, error)
}

type TransactionUsecase struct {
	trxManager      common.TransactionManager
	limitRepo       limit_repository.LimitRepositoryIF
	transactionRepo repository.TransactionRepositoryIF
}

func NewTransactionUsecase(trxManager common.TransactionManager, limitRepo limit_repository.LimitRepositoryIF, trxRepo repository.TransactionRepositoryIF) TransactionUsecaseIF {
	return &TransactionUsecase{
		trxManager:      trxManager,
		limitRepo:       limitRepo,
		transactionRepo: trxRepo,
	}
}

func (u *TransactionUsecase) CreateTransaction(ctx context.Context, body *dto.CreateTransaction, consumerId string) (string, error) {
	tx, err := u.trxManager.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer u.trxManager.Rollback(ctx, tx)

	limit, err := u.limitRepo.GetLimitWithLock(ctx, tx, consumerId, body.TenorMonth)

	if errors.Is(err, sql.ErrNoRows) {
		return "", util.NotFound("Limit tenor not found for this consumer")
	}

	if err != nil {
		return "", err
	}

	if limit.LimitAmount-limit.UsedAmount < body.OTRPrice {
		return "", fmt.Errorf("insufficient limit")
	}

	err = u.limitRepo.UpdateUsedAmountWithTx(ctx, tx, limit.ID, body.OTRPrice)
	if err != nil {
		return "", err
	}

	transaction := &model.Transaction{
		ID:             uuid.New().String(),
		ConsumerID:     consumerId,
		LimitID:        limit.ID,
		ContractNumber: fmt.Sprintf("TRX-%s", uuid.New().String()),
		OTRPrice:       body.OTRPrice,
		AdminFee:       body.AdminFee,
		Installment:    body.Installment,
		Interest:       body.Interest,
		AssetName:      body.AssetName,
		CreatedAt:      time.Now(),
	}

	err = u.transactionRepo.CreateTransactionWithTx(ctx, tx, transaction)
	if err != nil {
		return "", err
	}

	if err := u.trxManager.Commit(ctx, tx); err != nil {
		return "", err
	}

	return transaction.ContractNumber, nil
}

func (u *TransactionUsecase) GetTransactionHistory(ctx context.Context, params dto.GetTransactionHistoryQuery) ([]*dto.GetTransactionHistoryResponse, int, error) {
	offset := (params.Page - 1) * params.Limit

	result, err := u.transactionRepo.FetchTransactionByConsumer(ctx, params.ConsumerId, params.Limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := u.transactionRepo.CountTransactionByConsumer(ctx, params.ConsumerId)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
