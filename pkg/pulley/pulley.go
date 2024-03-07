package pulley

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Request(url string) []byte {
	resp, err := c.client.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Request failed: %s", err.Error()))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(fmt.Sprintf("Body parsing failed: %s", err.Error()))
	}

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Request failed with status %d\nBody: %s", resp.StatusCode, string(body)))
	}

	return body
}
