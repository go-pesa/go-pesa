package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-pesa/internal/globals"
	"github.com/go-pesa/internal/types"
	"github.com/google/go-querystring/query"
)

//Get is a http get
func Get(endpoint string, headers []types.Header, queryStrings interface{}) *http.Response {

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

	return res

}
