package gojek

import "net/http"

var API_URL = "https://api.gojekapi.com"

func Request(method string, url string) (*Response, error) {
	var getURL = API_URL + url

	client := &http.Client{}

	req, err := http.NewRequest(method, getURL, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-AppVersion", GetAppVersion())
	req.Header.Add("X-UniqueId", GetUniqueId())
	req.Header.Add("X-Location", GetLocation())
	req.Header.Add("Authorization", "Bearer "+GetToken())

	return client.Do(req)
}
