package entity

type Entity interface {
	EntityID() ID
	TableName() string
}
