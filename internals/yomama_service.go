package internals

import (
	"encoding/json"
	"net/http"
)

var baseURL = "https://yomama-jokes.com/api/random"

func createRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (app *Config) getYoMammaJoke() (*YoMamaJoke, error) {
	var joke YoMamaJoke

	client := &http.Client{}
	req, err := createRequest("GET", baseURL)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&joke); err != nil {
		return nil, err
	}

	return &joke, nil
}
