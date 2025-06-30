package repository

import (
	"errors"
)

var ErrNotFound = errors.New("entity not found")

//type AccountRepository[T entity.Account] interface {
//	CommonBehaviour[entity.Account]
//}

//type MessageRepository[T entity.Message] interface {
//	CommonBehaviour[entity.Message]
//}
