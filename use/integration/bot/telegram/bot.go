package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/bot"
)

type telegramBot struct {
	log *log.Logger
	api *tgbotapi.BotAPI
}

func New(log *log.Logger) (bot.Client, error) {
	apiKey := os.Getenv("TELEGRAM_API_TOKEN")
	if apiKey == "" {
		return nil, bot.ErrApiKeyEmpty
	}

	botAPI, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, err
	}

	return &telegramBot{log: log, api: botAPI}, nil
}
