package telegram

import (
	"github.com/ashkanamani/chatbot/internal/service"
	"gopkg.in/telebot.v4"
	"log/slog"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App, apiToken string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  apiToken,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("failed to connect to telegram servers", "error", err.Error())
		return nil, err
	}
	tg := &Telegram{
		App: app,
		bot: b,
	}

	tg.setupMiddlewares()

	tg.setupHandlers()

	return tg, nil
}

func (t *Telegram) Start() {
	slog.Info("telegram bot started successfully.")
	t.bot.Start()
}
