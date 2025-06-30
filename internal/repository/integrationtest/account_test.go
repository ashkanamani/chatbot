package integrationtest

import (
	"context"
	"fmt"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
	"github.com/ashkanamani/chatbot/internal/repository/postgres"
	"github.com/ashkanamani/chatbot/internal/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountRedisRepository_GetAndSave(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("%s:%s", "127.0.0.1", redisPort))
	defer redisClient.Close()
	assert.NoError(t, err)

	ctx := context.Background()
	arr := repository.NewAccountRedisRepository(redisClient)

	err = arr.Save(ctx, entity.Account{
		Id:        21,
		FirstName: "Nishtman",
	})
	assert.NoError(t, err)

	err = arr.Save(ctx, entity.Account{
		Id:        22,
		FirstName: "RadioNishtman",
	})
	assert.NoError(t, err)

	val, err := arr.Get(ctx, entity.NewID("account", "21"))
	assert.NoError(t, err)
	assert.Equal(t, "Nishtman", val.FirstName)
	assert.Equal(t, int64(21), val.Id)

	val, err = arr.Get(ctx, entity.NewID("account", "22"))
	assert.NoError(t, err)
	assert.Equal(t, "RadioNishtman", val.FirstName)
	assert.Equal(t, int64(22), val.Id)

	_, err = arr.Get(ctx, entity.NewID("account", 404))
	assert.ErrorIs(t, err, repository.ErrNotFound)
}

func TestAccountPostgresRepository_GetAndSave(t *testing.T) {
	conn, err := postgres.NewPostgresConnection(
		fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%s/postgres?sslmode=disable", postgresPort)
		)
	defer conn.Close(context.Background())
	assert.NoError(t, err)

	// continue
}
