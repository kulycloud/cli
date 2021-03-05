package client

import (
	"fmt"
	apiMapping "github.com/kulycloud/api-server/mapping"
	apiServer "github.com/kulycloud/api-server/server"
	"net/http"
)

func (c *Client) GetRoute(namespace string, name string) (*apiMapping.IncomingRoute, error) {
	var route = &apiMapping.IncomingRoute{}
	return route, c.ExecuteRequest(
		http.MethodGet,
		fmt.Sprintf("%s/route/%s", namespace, name),
		nil,
		route)
}

func (c *Client) GetRoutes(namespace string) ([]apiServer.RouteListElement, error) {
	var routes []apiServer.RouteListElement
	return routes, c.ExecuteRequest(
		http.MethodGet,
		fmt.Sprintf("%s/route", namespace),
		nil,
		&routes)
}
