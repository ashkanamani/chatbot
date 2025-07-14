package service

import (
	"context"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountService_RedisCreateOrUpdateWithUserExists(t *testing.T) {
	accRedisRepo := repository.NewMockAccountRepository(t)
	accPostgresRepo := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRedisRepo, accPostgresRepo)

	accRedisRepo.On("Get", mock.Anything, entity.NewID("account", 33)).Return(
		entity.Account{Id: 33, FirstName: "Ashkan"}, nil,
	).Once()

	accRedisRepo.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "RadioNishtman"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		Id:        33,
		FirstName: "RadioNishtman",
	})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, int64(33), newAcc.Id)
	assert.Equal(t, "RadioNishtman", newAcc.FirstName)
}
