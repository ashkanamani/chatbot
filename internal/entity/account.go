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
	PhoneNumber string    `json:"phone_number"`
	JoinedAt    time.Time `json:"joined_at"`
	IsActive    bool      `json:"is_active"`
	Blocked     bool      `json:"blocked"`
	LinkToken   string    `json:"link_token"`
}

func (a Account) EntityID() ID {
	return ID(fmt.Sprintf("account:%d", a.Id))
}
