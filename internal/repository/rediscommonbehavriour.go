package repository

import (
	"context"
	"github.com/ashkanamani/chatbot/internal/entity"
	"github.com/ashkanamani/chatbot/pkg/jsonhelper"
	"github.com/pkg/errors"
	"github.com/redis/rueidis"
	"log/slog"
)

var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity](client rueidis.Client) *RedisCommonBehaviour[T] {
	return &RedisCommonBehaviour[T]{client: client}
}

func (r *RedisCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		if errors.Is(err, rueidis.Nil) {
			return t, ErrNotFound
		}
		slog.Error("could not get from redis", "err", err, "id", id.String())
		return t, err
	}

	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r *RedisCommonBehaviour[T]) Save(ctx context.Context, ent T) error {
	cmd := r.client.B().JsonSet().Key(ent.EntityID().String()).
		Path("$").Value(string(jsonhelper.Encode(ent))).Build()
	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		slog.Error("could not save entity to redis", "err", err, "entity", ent)
		return err
	}
	return nil
}
