package prestashopApi

import (
	"crypto/tls"
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
	Resource       ResourceService
	Product        ProductService
}

func NewPrestaShop(appName, urlStr, apiKey string) *PrestaShop {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	ps := PrestaShop{
		AppName: appName,
		BaseURL: baseURL,
	}

	client := ps.newClient(apiKey)
	ps.Resource = newResourceService(client)

	ps.Product = ProductService{
		Client: client,
	}

	return &ps
}

func (ps *PrestaShop) newClient(apiKey string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := Client{
		client:  &http.Client{Transport: tr},
		version: apiVersion,
		baseURL: ps.BaseURL,
		apiKey:  apiKey,
	}

	return &client
}
