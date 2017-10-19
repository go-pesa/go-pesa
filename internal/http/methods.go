package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/go-pesa/go-pesa/internal/types"
	"github.com/google/go-querystring/query"
)

//Get is a http get
func Get(endpoint, baseURL string, headers []types.Header, queryStrings interface{}) *http.Response {

	var params string
	if queryStrings == nil {
		params = ""
	} else {

		q, err := query.Values(queryStrings)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to process query string, error:%s", err))
		}
		params = fmt.Sprintf("?%s", q.Encode())
	}

	url := fmt.Sprintf("%s%s%s", baseURL, endpoint, params)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create get request, error:%s", err))
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to send get request, error:%s", err))
	}

	return res

}

//Post is a http post
func Post(endpoint, baseURL string, headers []types.Header, payload interface{}) *http.Response {

	url := fmt.Sprintf("%s%s", baseURL, endpoint)
	encodedPayload := new(bytes.Buffer)
	json.NewEncoder(encodedPayload).Encode(payload)

	req, err := http.NewRequest("POST", url, encodedPayload)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create post request, error:%s", err))
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}
	log.Println(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to send post request, error:%s", err))
	}

	//TODO - Remove when entire API is covered, useful for debugging
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(res)
	log.Println(string(body))
	//

	return res

}
