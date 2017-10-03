package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

//Unmarshal unmarshals json data from API
func Unmarshal(res *http.Response, body interface{}) interface{} {
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Println(err)
	}
	return body
}
