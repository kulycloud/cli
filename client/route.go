package client

import (
	"encoding/json"
	"fmt"
	apiMapping "github.com/kulycloud/api-server/mapping"
	apiServer "github.com/kulycloud/api-server/server"
	"net/http"
)

func (c *Client) GetRoute(namespace string, name string) (*apiMapping.IncomingRoute, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/route/%s", c.instance, namespace, name))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		route := &apiMapping.IncomingRoute{}
		return route, json.NewDecoder(resp.Body).Decode(route)
	} else {
		e := &ErrorType{}
		err := json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", e.Error)
	}
}

func (c *Client) GetRoutes(namespace string) ([]apiServer.RouteListElement, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/route", c.instance, namespace))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var routes []apiServer.RouteListElement
		return routes, json.NewDecoder(resp.Body).Decode(&routes)
	} else {
		e := &ErrorType{}
		err := json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", e.Error)
	}
}
