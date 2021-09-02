package linebot

import (
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yona3/go-line-bot-sample/api"
)

func Webhook(events []*linebot.Event, r *http.Request) {
	handler, err := NewHandler()
	if err != nil {
		log.Println(err)
		return
	}
	bot, err := NewBot(handler)
	if err != nil {
		panic(err)
	}

	for _, e := range events {
		if e.Type == linebot.EventTypeMessage {
			switch msg := e.Message.(type) {
			case *linebot.TextMessage:
				text := msg.Text
				if strings.Contains(text, "いぬ") {
					replyDog := api.GetRandomDog()
					if _, err := bot.ReplyMessage(e.ReplyToken, linebot.NewImageMessage(replyDog, replyDog)).Do(); err != nil {
						log.Print(err)
						bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("わんちゃんは見つかりませんでした…。")).Do()
					}
				} else if strings.Contains(text, "ねこ") {
					replyCat := api.GetRandomCat()
					if _, err := bot.ReplyMessage(e.ReplyToken, linebot.NewImageMessage(replyCat, replyCat)).Do(); err != nil {
						log.Print(err)
						bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("ねこちゃんは見つかりませんでした…。")).Do()
					}
				} else {
					replyText := "「いぬ」or「ねこ」でメッセージを送ってね！"
					if _, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyText)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
