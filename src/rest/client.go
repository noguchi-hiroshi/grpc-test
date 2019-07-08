package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Get() (*Response, error)
}

type client struct {
}

func NewClient() Client {
	return &client{}
}

func (c *client) Get() (*Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := Response{}
	err = json.Unmarshal(body, &ret)
	return &ret, err
}
