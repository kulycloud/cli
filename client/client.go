package client

type Client struct {
	instance string
}

func NewClient(instance string) *Client {
	return &Client {
		instance: instance,
	}
}