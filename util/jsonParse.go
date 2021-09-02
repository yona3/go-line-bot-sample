package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func JsonParse(r *http.Response, d interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
