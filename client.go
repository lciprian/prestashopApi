package prestashopApi

import (
	"net/http"
	"net/url"
)

type Options struct {
	AppName        string
	CustomerKey    string
	CustomerSecret string
	Product        ProductService
	Client         *Client
}

type Client struct {
	client  *http.Client
	version string
	baseURL *url.URL
}
