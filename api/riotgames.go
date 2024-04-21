package riotgames

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"lol/common/dto"
	"net/http"
	"net/url"
)

const BASE_URL = "https://jp1.api.riotgames.com"

type ApiClient struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) *ApiClient {
	apiClient := &ApiClient{apiKey, &http.Client{}}
	return apiClient
}

func (c ApiClient) doRequest(method, apiUrl string, params map[string]string, data []byte) ([]byte, error) {
	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}
	endpoint := baseUrl.ResolveReference(u).String()

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err

	}
	q := req.URL.Query()
	q.Set("api_key", c.apiKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func (c ApiClient) GetChampionRotations() (*dto.ChampionInfo, error) {
	apiUrl := "/lol/platform/v3/champion-rotations"
	resp, err := c.doRequest("GET", apiUrl, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", apiUrl, string(resp))
	if err != nil {
		return nil, err
	}

	var championInfo dto.ChampionInfo
	if err := json.Unmarshal(resp, &championInfo); err != nil {
		return nil, err
	}
	return &championInfo, nil
}
