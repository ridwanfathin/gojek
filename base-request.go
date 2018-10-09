package gojek

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var API_URL = "https://api.gojekapi.com"

func Request(options map[string]string, url string) []byte {
	var getURL = API_URL + url

	req, _ := http.NewRequest(options["method"], getURL,
		bytes.NewBuffer([]byte(options["body"])))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-AppVersion", GetAppVersion())
	req.Header.Add("X-UniqueId", GetUniqueId())
	req.Header.Add("X-Location", GetLocation())
	req.Header.Add("Authorization", "Bearer "+GetToken())

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf(fmt.Sprint(err))
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// var response = map[string]string{}
	// json.Unmarshal(body, &response)

	// fmt.Println(response)

	return body
}
