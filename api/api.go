package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yona3/go-line-bot-sample/util"
)

type DogData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type CatData struct {
	File string `json:"file"`
}

func GetRandomCat() {
	url := os.Getenv("RANDOM_CAT_URL")
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var d CatData
	err = util.JsonParse(res, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", d.File)
}

func GetRandomDog() {
	url := os.Getenv("RANDOM_DOG_URL")
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var d DogData
	err = util.JsonParse(res, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", d.Message)
}
