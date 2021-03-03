package client

import (
	"encoding/json"
	"fmt"
	apiServer "github.com/kulycloud/api-server/server"
	protoStorage "github.com/kulycloud/protocol/storage"
	"net/http"
)

type ErrorType struct {
	Error string `json:"error"`
}

func (c *Client) GetService(namespace string, name string) (*protoStorage.Service, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/service/%s", c.instance, namespace, name))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		svc := &protoStorage.Service{}
		return svc, json.NewDecoder(resp.Body).Decode(svc)
	} else {
		e := &ErrorType{}
		err := json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", e.Error)
	}
}

func (c *Client) GetServices(namespace string) ([]apiServer.ServiceListElement, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/service", c.instance, namespace))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var svc []apiServer.ServiceListElement
		return svc, json.NewDecoder(resp.Body).Decode(&svc)
	} else {
		e := &ErrorType{}
		err := json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", e.Error)
	}
}
