package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CatData struct {
	File string `json:"file"`
}

func GetRandomCat() {
	url := os.Getenv("RANDOM_CAT_URL")
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var d CatData
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", d.File)
}
