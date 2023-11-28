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
var productVariantBasePath = "combinations"

type ProductService struct {
	client *Client
}

type ProductList struct {
	Limit int
	Page  int
	Data  []models.Product `json:"products,omitempty"`
}

type ProductCombinationList struct {
	Limit int
	Page  int
	Data  []models.Combination `json:"combinations,omitempty"`
}

func newProductService(client *Client) ProductService {
	return ProductService{
		client: client,
	}
}

func (s *ProductService) ListProducts(limit, page int) (*ProductList, error) {
	productList := ProductList{
		Limit: limit,
		Page:  page,
	}

	if page > 0 {
		page -= 1
	}
	offset := limit * page
	queryParams := url.Values{}
	queryParams.Add("display", "full")
	queryParams.Add("limit", fmt.Sprintf("%d,%d", offset, limit))

	//products := make([]models.Product, 0)
	if err := s.client.Get(productBasePath, queryParams, &productList); err != nil {
		return nil, err
	}
	//	fmt.Println("-ListProducts---------", products)
	return &productList, nil
}

func (s *ProductService) CreateProduct(product models.ProductRequest) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}

func (s *ProductService) ListProductCombination(productId string, limit, page int) (*ProductCombinationList, error) {
	if page > 0 {
		page -= 1
	}
	offset := limit * page
	queryParams := url.Values{}
	queryParams.Add("display", "full")
	queryParams.Add("limit", fmt.Sprintf("%d,%d", offset, limit))
	if productId != "" {
		queryParams.Add("filter[id_product]", productId)
	}

	productCombinationList := ProductCombinationList{
		Limit: limit,
		Page:  page,
	}

	if err := s.client.Get(productVariantBasePath, queryParams, &productCombinationList); err != nil {
		return nil, err
	}
	fmt.Println("--", productCombinationList.Data)
	panic("----------")
	return &productCombinationList, nil
}

func (s *ProductService) CreateProductCombination(productVariant models.ProductVariantReq) (*models.Combination, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Combinations: &productVariant})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productVariantBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Combination, nil
}
