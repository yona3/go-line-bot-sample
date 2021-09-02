package main

import (
	"fmt"

	"github.com/yona3/go-line-bot-sample/api"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	api.GetRandomCat()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
