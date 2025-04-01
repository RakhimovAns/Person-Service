package client

import (
	"encoding/json"
	"fmt"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"io"
	"net/http"
)

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

type NationalizeClient interface {
	GetNationality(name string) (string, error)
}

type nationalizeClient struct {
	baseURL string
	logger  logging.Logger
}

func NewNationalizeClient(baseURL string, logger logging.Logger) NationalizeClient {
	return &nationalizeClient{
		baseURL: baseURL,
		logger:  logger,
	}
}

func (c *nationalizeClient) GetNationality(name string) (string, error) {
	url := fmt.Sprintf("%s/?name=%s", c.baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		c.logger.Error("Failed to make request to Nationalize API: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read Nationalize API response: %v", err)
		return "", err
	}

	var nationalizeResponse NationalizeResponse
	if err := json.Unmarshal(body, &nationalizeResponse); err != nil {
		c.logger.Error("Failed to unmarshal Nationalize API response: %v", err)
		return "", err
	}

	if len(nationalizeResponse.Country) == 0 {
		c.logger.Debug("No country data for name %s", name)
		return "", nil
	}

	c.logger.Debug("Received country %s for name %s", nationalizeResponse.Country[0].CountryID, name)
	return nationalizeResponse.Country[0].CountryID, nil
}
