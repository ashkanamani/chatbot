package repository

import (
	"context"
	"errors"
	"github.com/ashkanamani/chatbot/internal/entity"
)

var ErrNotFound = errors.New("entity not found")

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, entity T) error
}

type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

//type AccountRepository[T entity.Account] interface {
//	CommonBehaviour[entity.Account]
//}

//type MessageRepository[T entity.Message] interface {
//	CommonBehaviour[entity.Message]
//}
