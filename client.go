package harper

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	endpoint   string
	HttpClient *resty.Client
}

func NewClientWithHTTPClient(httpClient *http.Client, endpoint, username, password string) *Client {
	restyClient := resty.
		NewWithClient(httpClient).
		SetDisableWarn(true).
		SetBasicAuth(username, password)
	
	return &Client{
		endpoint:   endpoint,
		HttpClient: restyClient,
	}
}

func NewClient(endpoint, username, password string) *Client {
	httpClient := resty.
		New().
		SetDisableWarn(true).
		SetBasicAuth(username, password)

	return &Client{
		endpoint:   endpoint,
		HttpClient: httpClient,
	}
}

// RawRequest allows raw requests to be made against the client.
// The recommended route for making calls is via the specific function endpoints.
func (c *Client) RawRequest(op Operation, result interface{}) error {
	return c.opRequest(op, result)
}

func (c *Client) opRequest(op Operation, result interface{}) error {
	e := ErrorResponse{}

	req := c.HttpClient.
		NewRequest().
		SetBody(op.Prepare()).
		SetError(&e)

	if result != nil {
		req.SetResult(result)
	}

	resp, err := req.Post(c.endpoint)
	if err != nil {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    err.Error()}
	}
	if resp.StatusCode() > 399 {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    string(resp.Body())}
	}

	return nil
}

func (c *Client) SetConfigurationRequest(op interface{}, result interface{}) error {
	e := ErrorResponse{}
	req := c.HttpClient.NewRequest().SetBody(op).SetError(&e)

	if result != nil {
		req.SetResult(result)
	}

	resp, err := req.Post(c.endpoint)
	if err != nil {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    err.Error()}
	}
	if resp.StatusCode() > 399 {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    string(resp.Body())}
	}

	return nil
}

// Healthcheck does a GET request against the /health endpoint of the
// configured Harper server and returns an error if it gets a non-200
// status, nil otherwise.
func (c *Client) Healthcheck() error {
	e := ErrorResponse{}

	healthCheckURL, err := url.JoinPath(c.endpoint, "health")
	if err != nil {
		return fmt.Errorf("invalid healthcheck URL: '%w'", err)
	}

	req := c.HttpClient.NewRequest().SetError(&e)

	resp, err := req.Get(healthCheckURL)
	if err != nil {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    err.Error(),
		}
	}
	if resp.StatusCode() > 399 {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    string(resp.Body()),
		}
	}

	return nil
}
