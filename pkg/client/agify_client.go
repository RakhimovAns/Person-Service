package client

import (
	"encoding/json"
	"fmt"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"io"
	"net/http"
)

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type AgifyClient interface {
	GetAge(name string) (int, error)
}

type agifyClient struct {
	baseURL string
	logger  logging.Logger
}

func NewAgifyClient(baseURL string, logger logging.Logger) AgifyClient {
	return &agifyClient{
		baseURL: baseURL,
		logger:  logger,
	}
}

func (c *agifyClient) GetAge(name string) (int, error) {
	url := fmt.Sprintf("%s/?name=%s", c.baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		c.logger.Error("Failed to make request to Agify API: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read Agify API response: %v", err)
		return 0, err
	}

	var agifyResponse AgifyResponse
	if err := json.Unmarshal(body, &agifyResponse); err != nil {
		c.logger.Error("Failed to unmarshal Agify API response: %v", err)
		return 0, err
	}

	c.logger.Debug("Received age %d for name %s", agifyResponse.Age, name)
	return agifyResponse.Age, nil
}
