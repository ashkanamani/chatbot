package entity

import (
	"fmt"
	"time"
)

type Account struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	JoinedAt    time.Time `json:"joined_at"`
	IsActive    bool      `json:"is_active"`
	Blocked     bool      `json:"blocked"`
	LinkToken   string    `json:"link_token"`
	State       string    `json:"state"`
}

func (a Account) TableName() string {
	return "accounts"
}

func (a Account) EntityID() ID {
	return ID(fmt.Sprintf("account:%d", a.Id))
}
