package service

import (
	"context"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
)

type AccountService struct {
	redisAccounts    repository.AccountRepository
	postgresAccounts repository.AccountRepository
}

func NewAccountService(redisAccounts repository.AccountRepository, postgresAccount repository.AccountRepository) *AccountService {
	return &AccountService{
		redisAccounts:    redisAccounts,
		postgresAccounts: postgresAccount,
	}
}

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) error {
	return a.redisAccounts.Save(ctx, account)
}
