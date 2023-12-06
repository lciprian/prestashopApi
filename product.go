package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productBasePath = "products"

type ProductList struct {
	Limit int
	Page  int
	Data  []models.ProductL `json:"products,omitempty"`
}

type ProductService struct {
	client *Client
}

func newProductService(client *Client) ProductService {
	return ProductService{
		client: client,
	}
}

func (s *ProductService) ListProducts(limit, page int) (*ProductList, error) {
	if page > 0 {
		page -= 1
	}
	offset := limit * page
	queryParams := url.Values{}
	queryParams.Add("display", "full")
	queryParams.Add("limit", fmt.Sprintf("%d,%d", offset, limit))

	productList := ProductList{Limit: limit, Page: page}
	if err := s.client.Get(productBasePath, queryParams, &productList); err != nil {
		return nil, err
	}

	return &productList, nil
}

func (s *ProductService) CreateProduct(product models.ProductReq) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}

func (s *ProductService) UpdateProduct(product models.ProductReq) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Put(productBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}
