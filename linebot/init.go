package linebot

import (
	"os"

	sdk "github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func NewHandler() (*httphandler.WebhookHandler, error) {
	secret := os.Getenv("LINE_CHANNEL_SECRET")
	token := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	handler, err := httphandler.New(secret, token)
	return handler, err
}

func NewBot(handler *httphandler.WebhookHandler) (*sdk.Client, error) {
	bot, err := handler.NewClient()
	return bot, err
}
