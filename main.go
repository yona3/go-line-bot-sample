package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yona3/go-line-bot-sample/linebot"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	handler, err := linebot.NewHandler()
	if err != nil {
		log.Fatal(err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	handler.HandleEvents(linebot.Webhook)
	http.Handle("/callback", handler)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
