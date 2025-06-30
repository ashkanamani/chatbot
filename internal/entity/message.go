package entity

import (
	"fmt"
	"time"
)

type Message struct {
	MessageId        int64     `json:"message_id"`
	FromUserId       int64     `json:"from_user_id"`
	ToUserId         int64     `json:"to_user_id"`
	Text             string    `json:"text"`
	Date             time.Time `json:"date"`
	Delivered        bool      `json:"delivered"`
	ReplyToMessageId int64     `json:"reply_to_message_id"`
}

func (m Message) EntityID() ID {
	return ID(fmt.Sprintf("message:%d", m.MessageId))
}

func (m Message) TableName() string {
	return "messages"
}
