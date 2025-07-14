package telegram

import (
	"context"
	"github.com/ashkanamani/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) setupMiddlewares() {
	t.bot.Use(t.registerMiddleware)
}

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		_ = c.Reply("سڵامەلەی")
		acc := entity.Account{
			Id:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			LastName:  c.Sender().LastName,
			Username:  c.Sender().Username,
			IsActive:  true,
			Blocked:   false,
		}
		acc, created, err := t.App.Accounts.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			return err
		}
		c.Set("account", acc)
		c.Set("is_just_created", created)
		return next(c)
	}
}
