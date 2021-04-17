package joke

import (
	"encoding/json"
	"fmt"
	"net/http"
	"gitlab.com/Zendden/workshop/internal/api"
)

const apiPath = "/api?format=json"

type JockClient struct {
	url string
}

func NewClient(baseUrl string) *JockClient {
	return &JockClient{
		url: baseUrl,
	}
}

func (jClient *JockClient) Get() (*api.JokeResponse, error) {
	urlPath := jClient.url + apiPath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}