package client

import (
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
		fmt.Sprintf("%s/services", namespace),
		nil,
		&services)
}
