package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](
	fullUrl string,
	httpMethod string,
	headers map[string]string,
	queryParameters url.Values,
	body io.Reader,
	responseType T, 
	) (T, error) {
		client := http.Client{}
		parsedUrl, err := url.Parse(fullUrl)
		if err != nil {
			return responseType, err
		}

		// append query params
		if httpMethod == "GET" {
			query := parsedUrl.Query()

			for k, v := range queryParameters {
				query.Set(k, strings.Join(v, ","))
			}

			parsedUrl.RawQuery = query.Encode()
		}

		req, err := http.NewRequest(httpMethod, parsedUrl.String(), body)
		if err != nil {
			return responseType, err
		}

		// parse headers and add value to request
		for k, v := range headers {
			req.Header.Set(k, v)
		}

		// optional
		// log.Printf("%s %s\n", httpMethod, req.URL.String())

		// execute request
		res, err := client.Do(req)
		if err != nil {
			return responseType, err
		}

		if res == nil {
			return responseType, fmt.Errorf("error: calling %s returned empty response", parsedUrl.String())
		}

		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			return responseType, err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return responseType, fmt.Errorf("error calling %s:\nstatus: %s\nresponse: %s", parsedUrl.String(), res.Status, responseData)
		}

		var responseObject T
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			// log.Printf("error unmarshaling response: %+v", err)
			return responseType, err
		}

		return responseObject, nil
}

func main() {
	url := "https://api.github.com/users/rogerwaldron"
	headers := map[string]string {
		"Accept": "application/json",
	}
	queryParameters := url.Values{}
	var response map[string]interface{}
	response, err := MakeHTTPRequest(url, "GET", headers, queryParameters, nil, response)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %+v", response)
}