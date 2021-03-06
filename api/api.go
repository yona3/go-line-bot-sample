package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yona3/go-line-bot-sample/util"
)

type CatResponse struct {
	File string `json:"file"`
}

func GetRandomCat() string {
	url := os.Getenv("RANDOM_CAT_URL")
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var d CatResponse
	if err = util.JsonParse(res, &d); err != nil {
		log.Fatal(err)
	}

	return d.File
}

type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetRandomDog() string {
	url := os.Getenv("RANDOM_DOG_URL")
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var d DogResponse
	if err = util.JsonParse(res, &d); err != nil {
		log.Fatal(err)
	}

	return d.Message
}
