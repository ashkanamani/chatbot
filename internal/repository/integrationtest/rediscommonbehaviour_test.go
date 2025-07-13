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

type TestType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t TestType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func (t TestType) TableName() string {
	return "test"
}

func TestRedisCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	defer redisClient.Close()
	ctx := context.Background()

	rcb := repository.NewRedisCommonBehaviour[TestType](redisClient)
	err = rcb.Save(ctx, TestType{
		ID:   "11",
		Name: "Sheldon",
	})
	assert.NoError(t, err)

	val, err := rcb.Get(ctx, entity.NewID("testType", "11"))
	assert.NoError(t, err)
	assert.Equal(t, "Sheldon", val.Name)
	assert.Equal(t, "11", val.ID)

	_, err = rcb.Get(ctx, entity.NewID("testType", "12"))
	assert.ErrorIs(t, err, repository.ErrNotFound)
}
