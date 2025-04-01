package client

import (
	"encoding/json"
	"fmt"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"io"
	"net/http"
)

type GenderizeResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type GenderizeClient interface {
	GetGender(name string) (string, error)
}

type genderizeClient struct {
	baseURL string
	logger  logging.Logger
}

func NewGenderizeClient(baseURL string, logger logging.Logger) GenderizeClient {
	return &genderizeClient{
		baseURL: baseURL,
		logger:  logger,
	}
}

func (c *genderizeClient) GetGender(name string) (string, error) {
	url := fmt.Sprintf("%s/?name=%s", c.baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		c.logger.Error("Failed to make request to Genderize API: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read Genderize API response: %v", err)
		return "", err
	}

	var genderizeResponse GenderizeResponse
	if err := json.Unmarshal(body, &genderizeResponse); err != nil {
		c.logger.Error("Failed to unmarshal Genderize API response: %v", err)
		return "", err
	}

	c.logger.Debug("Received gender %s for name %s", genderizeResponse.Gender, name)
	return genderizeResponse.Gender, nil
}
