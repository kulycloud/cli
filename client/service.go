package client

import (
	"bytes"
	"fmt"
	apiServer "github.com/kulycloud/api-server/server"
	protoStorage "github.com/kulycloud/protocol/storage"
	"net/http"
)

type ErrorType struct {
	Error string `json:"error"`
}

func (c *Client) GetService(namespace string, name string) (*protoStorage.Service, error) {
	var svc = &protoStorage.Service{}
	return svc, c.ExecuteRequest(
		http.MethodGet,
		fmt.Sprintf("%s/service/%s", namespace, name),
		nil,
		svc)
}

func (c *Client) GetServices(namespace string) ([]apiServer.ServiceListElement, error) {
	var services []apiServer.ServiceListElement
	return services, c.ExecuteRequest(
		http.MethodGet,
		fmt.Sprintf("%s/service", namespace),
		nil,
		&services)
}

func (c *Client) CreateService(namespace, name string, config []byte) (*protoStorage.Service, error) {
	var service = &protoStorage.Service{}
	return service, c.ExecuteRequest(
		http.MethodPost,
		fmt.Sprintf("%s/service/%s", namespace, name),
		bytes.NewReader(config),
		service)
}

func (c *Client) UpdateService(namespace, name string, config []byte) (*protoStorage.Service, error) {
	var service = &protoStorage.Service{}
	return service, c.ExecuteRequest(
		http.MethodPut,
		fmt.Sprintf("%s/service/%s", namespace, name),
		bytes.NewReader(config),
		service)
}

func (c *Client) DeleteService(namespace, name string) (*protoStorage.Service, error) {
	var service = &protoStorage.Service{}
	return service, c.ExecuteRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/service/%s", namespace, name),
		nil,
		service)
}
