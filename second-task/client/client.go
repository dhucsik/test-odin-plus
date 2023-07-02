package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"second-task/model"
	"strings"
)

type DefaulClient struct {
	HTTPClient *http.Client
}

func New() *DefaulClient {
	return &DefaulClient{
		HTTPClient: &http.Client{},
	}
}

func (c *DefaulClient) GetRequest(currencies *[]model.Currency, url string) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}

	if err := json.Unmarshal(responseBody, currencies); err != nil {
		return fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
	}

	return nil
}

func (c *DefaulClient) FilterByName(currencies []model.Currency, searchName string) []model.Currency {
	out := make([]model.Currency, 0)

	for _, currency := range currencies {
		if strings.Contains(currency.Name, searchName) {
			out = append(out, currency)
		}
	}

	return out
}
