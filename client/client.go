package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sm-manager/model"
	"strings"
)

type Client struct {
	basePath string
}

func NewClient(basePath string) *Client {
	return &Client{
		basePath: basePath,
	}
}

func (c *Client) buildURL(endpoint string) string {
	if !strings.HasSuffix(c.basePath, "/") {
		c.basePath += "/"
	}

	return fmt.Sprintf("%s%s", c.basePath, endpoint)
}

func (c *Client) GetStatus() (*model.StatusResponse, error) {
	request, err := http.NewRequest(http.MethodGet, c.buildURL("api/v1/status"), nil)
	if err != nil {
		return nil, err
	}

	bodyRaw, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var statusResponse model.StatusResponse

	err = json.Unmarshal(bodyRaw, &statusResponse)
	if err != nil {
		return nil, err
	}

	return &statusResponse, nil
}
