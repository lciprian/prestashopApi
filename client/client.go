package client

import (
	"net/http"
)

type Options struct {
	BaseAPIURL string

	// Basic Auth
	Username string
	Password string
}

type Client struct {
	baseURL string
	req     *http.Request
	options *Options
}

func NewClient(options *Options) *Client {
	req := http.Request{}

	return &Client{
		req:     &req,
		options: options,
		baseURL: options.BaseAPIURL,
	}
}
