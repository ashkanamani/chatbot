package integrationtest

import (
	"context"
	"fmt"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/internal/repository"
	"github.com/ashkanamani/chatbot/internal/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("%s:%s", "127.0.0.1", redisPort))
	defer redisClient.Close()
	assert.NoError(t, err)

	ctx := context.Background()
	rcb := repository.NewRedisCommonBehaviour[testType](redisClient)

	err = rcb.Save(ctx, testType{
		ID:   "21",
		Name: "Nishtman",
	})
	assert.NoError(t, err)

	err = rcb.Save(ctx, testType{
		ID:   "22",
		Name: "RadioNishtman",
	})
	assert.NoError(t, err)

	val, err := rcb.Get(ctx, entity.NewID("testType", "21"))
	assert.NoError(t, err)
	assert.Equal(t, "Nishtman", val.Name)
	assert.Equal(t, "21", val.ID)

	val, err = rcb.Get(ctx, entity.NewID("testType", "22"))
	assert.NoError(t, err)
	assert.Equal(t, "RadioNishtman", val.Name)
	assert.Equal(t, "22", val.ID)

	_, err = rcb.Get(ctx, entity.NewID("testType", "404"))
	assert.ErrorIs(t, err, repository.ErrNotFound)
}
