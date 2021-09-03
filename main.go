package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sdk "github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yona3/go-line-bot-sample/linebot"
)

func main() {
	handler, err := linebot.NewHandler()
	if err != nil {
		log.Fatal(err)
		return
	}

	bot, err := linebot.NewBot(handler)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	handler.HandleEvents(func(events []*sdk.Event, r *http.Request) {
		linebot.Webhook(bot, events, r)
	})
	http.Handle("/callback", handler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
