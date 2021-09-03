package linebot

import (
	"log"
	"net/http"
	"strings"

	sdk "github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yona3/go-line-bot-sample/api"
)

func Webhook(bot *sdk.Client, events []*sdk.Event, r *http.Request) {
	for _, e := range events {
		if e.Type == sdk.EventTypeMessage {
			switch msg := e.Message.(type) {
			case *sdk.TextMessage:
				text := msg.Text
				if strings.Contains(text, "いぬ") {
					replyDog := api.GetRandomDog()
					if _, err := bot.ReplyMessage(e.ReplyToken, sdk.NewImageMessage(replyDog, replyDog)).Do(); err != nil {
						log.Print(err)
						bot.ReplyMessage(e.ReplyToken, sdk.NewTextMessage("わんちゃんは見つかりませんでした…。")).Do()
					}
				} else if strings.Contains(text, "ねこ") {
					replyCat := api.GetRandomCat()
					if _, err := bot.ReplyMessage(e.ReplyToken, sdk.NewImageMessage(replyCat, replyCat)).Do(); err != nil {
						log.Print(err)
						bot.ReplyMessage(e.ReplyToken, sdk.NewTextMessage("ねこちゃんは見つかりませんでした…。")).Do()
					}
				} else {
					replyText := "「いぬ」or「ねこ」でメッセージを送ってね！"
					if _, err := bot.ReplyMessage(e.ReplyToken, sdk.NewTextMessage(replyText)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
