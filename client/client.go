package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	instance string
}

func NewClient(instance string) *Client {
	return &Client {
		instance: instance,
	}
}

func (c *Client) ExecuteRequest(method, url string, body io.Reader, resource interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.instance, url), body)
	if err != nil {
		return fmt.Errorf("request preparation failed: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(resource)
	} else {
		e := &ErrorType{}
		bodyRaw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("server replied with error but reading failed: %w", err)
		}
		err = json.Unmarshal(bodyRaw, e)
		if err != nil {
			return fmt.Errorf("server replied with error but decoding failed. Raw error: %s",
				strings.TrimSpace(string(bodyRaw)))
		}
		return fmt.Errorf("server replied with error: %s", e.Error)
	}
}
