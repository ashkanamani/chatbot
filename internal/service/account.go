package service

import (
	"context"
	"errors"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
	"time"
)

const (
	DefaultState = "home"
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

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.redisAccounts.Get(ctx, account.EntityID())
	// User exists in the database
	if err == nil {
		if savedAccount.Username != account.Username ||
			savedAccount.FirstName != account.FirstName ||
			savedAccount.DisplayName != account.DisplayName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			savedAccount.DisplayName = account.DisplayName
			return savedAccount, false, a.redisAccounts.Save(ctx, savedAccount)
		}
	}
	// User does not exist in the database
	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.redisAccounts.Save(ctx, account)
	}
	return savedAccount, false, err
}
