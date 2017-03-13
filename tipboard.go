package gotipboard

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Config used by the client tipboard for communicate with the API
type Config struct {
	BaseURL    string
	APIVersion string
	APIKey     string
}

type Client struct {
	config Config
	client http.Client
}

// NewClient return client for use API of tipboard
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		client: http.Client{},
	}
}

// Push send data to tipboard API
func (tipboard *Client) Push(tile string, key string, data interface{}) error {
	values, err := json.Marshal(data)

	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("tile", tile)
	form.Add("key", key)
	form.Add("data", string(values))

	request, err := http.NewRequest(http.MethodPost, tipboard.getUrlPush(), strings.NewReader(form.Encode()))

	if err != nil {
		return err
	}

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New(response.Status) // TODO improve error
	}

	return nil
}

func (tipboard *Client) getUrlPush() string {
	return fmt.Sprintf("%s/%s/%s/push", tipboard.config.BaseURL, tipboard.config.APIVersion, tipboard.config.APIKey)
}
