package adapters

import (
	"encoding/json"
	"golang-test/internal/application/dtos"
	"net/http"
)

type BoostrResponse struct {
	Status string            `json:"status"`
	Data   []dtos.HolidayDto `json:"data"`
}

type BoostrClient struct {
	url string
}

func NewBoostrClient(url string) *BoostrClient {
	return &BoostrClient{url: url}
}

func (b *BoostrClient) FetchHolidays() ([]dtos.HolidayDto, error) {
	response, err := http.Get(b.url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var parsed BoostrResponse
	if err := json.NewDecoder(response.Body).Decode(&parsed); err != nil {
		return nil, err
	}
	return parsed.Data, nil
}
