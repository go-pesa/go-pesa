package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/types"
	"github.com/google/go-querystring/query"
)

func get(endpoint string, headers []types.Header, queryStrings interface{}) []byte {

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

	url := fmt.Sprintf("%s%s%s", globals.BaseURL, endpoint, params)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create get request, error:%s", err))
	}
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to send get request, error:%s", err))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to process response body, error:%s", err))
	}
	defer res.Body.Close()

	return body

}
