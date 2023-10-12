package prestashopApi

import (
	"fmt"
	"net/url"
)

var productBasePath = "products"

type ProductService struct {
	client *Client
}

type Product struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductList struct {
	Products []Product `json:"products,omitempty"`
}

func newProductService(client *Client) ProductService {
	return ProductService{
		client: client,
	}
}

func (s *ProductService) ListProducts(limit, page int) ([]Product, error) {
	productList := ProductList{}

	if page > 0 {
		page -= 1
	}
	offset := limit * page
	queryParams := url.Values{}
	queryParams.Add("display", "full")
	queryParams.Add("limit", fmt.Sprintf("%d,%d", offset, limit))

	err := s.client.Get(productBasePath, queryParams, &productList)
	if err != nil {
		return nil, err
	}

	return productList.Products, nil
}
