package integrationtest

import (
	"context"
	"fmt"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
	"github.com/ashkanamani/chatbot/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountPostgresRepositoryGetAndSet(t *testing.T) {
	conn, err := postgres.NewPostgresConnection(
		fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%s/postgres?sslmode=disable", postgresPort),
	)
	assert.NoError(t, err)
	par := repository.NewAccountPostgresRepository(conn, "accounts")

	ctx := context.Background()
	iranLoc, _ := time.LoadLocation("Asia/Tehran")
	err = par.Save(ctx, entity.Account{
		Id:          int64(21),
		FirstName:   "Nishtman",
		LastName:    "Kurdi",
		Username:    "RadioNishtman",
		DisplayName: "Kurd",
		JoinedAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, iranLoc),
		IsActive:    true,
		Blocked:     false,
		LinkToken:   "https://t.me/radionishtman",
	})
	assert.NoError(t, err)

	val, err := par.Get(ctx, entity.NewID("account", int64(21)))

	assert.NoError(t, err)
	assert.Equal(t, "Nishtman", val.FirstName)
	assert.Equal(t, "Kurdi", val.LastName)
	assert.Equal(t, "RadioNishtman", val.Username)
	assert.Equal(t, "Kurd", val.DisplayName)
	assert.True(t, val.JoinedAt.Equal(time.Date(2000, 1, 1, 0, 0, 0, 0, iranLoc)))
	assert.Equal(t, true, val.IsActive)
	assert.Equal(t, false, val.Blocked)
	assert.Equal(t, "https://t.me/radionishtman", val.LinkToken)
}
