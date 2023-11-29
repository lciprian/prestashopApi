package prestashopApi

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

var (
	apiVersion = "1.7.8.8"
)

const (
	UserAgent            = "prestashopApi/1.0.0"
	defaultHttpTimeout   = 30
	defaultApiPathPrefix = "/api"
	defaultVersion       = "1.7.8.8"
)

type PrestaShop struct {
	AppName            string
	BaseURL            *url.URL
	Resource           ResourceService
	Product            ProductService
	ProductVariant     ProductVariantService
	ProductOption      ProductOptionService
	ProductOptionValue ProductOptionValueService
	Image              ImageService
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
	ps.Product = newProductService(client)
	ps.ProductVariant = newProductVariantService(client)
	ps.ProductOption = newProductOptionService(client)
	ps.ProductOptionValue = newProductOptionValueService(client)
	ps.Image = newImageService(client)

	return &ps
}

func (ps *PrestaShop) newClient(apiKey string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := Client{
		version: apiVersion,
		baseURL: ps.BaseURL,
		apiKey:  apiKey,
		client: &http.Client{
			Transport: tr,
			Timeout:   time.Second * defaultHttpTimeout,
		},
	}

	return &client
}
