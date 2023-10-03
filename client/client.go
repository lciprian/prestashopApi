package client

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	UserAgent            = "woocommerce/1.0.0"
	defaultHttpTimeout   = 10
	defaultApiPathPrefix = "/wp-json/wc/v3"
	defaultVersion       = "v3"
)

var (
	apiVersionRegex = regexp.MustCompile(`^v[0-9]{2}$`)
)

type App struct {
	CustomerKey    string
	CustomerSecret string
	AppName        string
	UserId         string
	Scope          string
	ReturnUrl      string
	CallbackUrl    string
	Client         *Client
}

type Options struct {
	AppName        string
	CustomerKey    string
	CustomerSecret string
	Client         *Client
}

type Client struct {
	Client     *http.Client
	app        App
	version    string
	baseURL    *url.URL
	pathPrefix string
	token      string

	//Product        ProductService
	//ProductVariant ProductVariantService
}

func (a App) NewClient(shopName string) *Client {
	return NewClient(a, shopName)
}

func NewClient(app App, shopName string) *Client {
	baseURL, err := url.Parse(ShopBaseURL(shopName))
	if err != nil {
		panic(err)
	}
	c := &Client{
		Client: &http.Client{
			Timeout: time.Second * defaultHttpTimeout,
		},
		app:        app,
		baseURL:    baseURL,
		version:    defaultVersion,
		pathPrefix: defaultApiPathPrefix,
	}
	//c.Product = &ProductServiceOp{client: c}
	//c.ProductVariant = &ProductVariantServiceOp{client: c}
	//c.Order = &OrderServiceOp{client: c}
	//c.OrderNote = &OrderNoteServiceOp{client: c}
	//c.Webhook = &WebhookServiceOp{client: c}
	//for _, opt := range opts {
	//	opt(c)
	//}

	return c
}

// ShopBaseURL return a shop's base https base url
func ShopBaseURL(shopName string) string {
	return fmt.Sprintf("https://%s", shopName)
}
