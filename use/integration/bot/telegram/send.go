package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/igilgyrg/arbitrage/use/integration/bot"
)

func (t *telegramBot) Send(ctx context.Context, message bot.Message) (err error) {
	botMessage := tgbotapi.NewMessage(message.ChatId, message.Content)
	send, err := t.api.Send(botMessage)
	if err != nil {
		return
	}

	if send.MessageID == 0 {
		err = bot.ErrSendMessage

		return
	}

	return
}
