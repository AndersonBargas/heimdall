package consumption

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Consumer struct {
	Client *http.Client
}

func NewConsumer(httpClient *http.Client) *Consumer {
	return &Consumer{
		Client: httpClient,
	}
}

func (c *Consumer) Consume(endpoint *url.URL) (int, []byte, error) {
	resp, err := c.Client.Get(endpoint.String())
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}
