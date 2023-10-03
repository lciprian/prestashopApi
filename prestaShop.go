package prestashopApi

import (
	"net/http"
	"net/url"
)

var (
	apiVersion = "1.7.8.8"
)

const (
	UserAgent            = "prestashopApi/1.0.0"
	defaultHttpTimeout   = 10
	defaultApiPathPrefix = "/api"
	defaultVersion       = "1.7.8.8"
)

type PrestaShop struct {
	AppName        string
	BaseURL        *url.URL
	apiKey         string
	CustomerSecret string
	Product        ProductService
}

func NewPrestaShop(appName, urlStr, customerKey string) *PrestaShop {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	ps := PrestaShop{
		AppName:        appName,
		BaseURL:        baseURL,
		apiKey:         customerKey,
		CustomerSecret: customerKey,
	}

	client := ps.newClient()
	ps.Product = ProductService{
		Client: client,
	}

	return &ps
}

func (ps PrestaShop) newClient() *Client {
	client := Client{
		client:  &http.Client{},
		version: apiVersion,
		baseURL: ps.BaseURL,
	}

	return &client
}
