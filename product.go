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

type ProductService struct {
	client *Client
}

type ProductList struct {
	Limit int
	Page  int
	Data  []models.Product `json:"products,omitempty"`
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

func (s *ProductService) CreateProduct(product models.ProductRequest) error {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Product: &product})
	if err != nil {
		return err
	}

	data, err := s.client.Post(productBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return err
	}

	fmt.Println("---------", psResponse.Product)
	return nil
}

func (s *ProductService) CreateProductCombination(productVariant models.ProductVariantReq) error {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Combinations: &productVariant})
	if err != nil {
		return err
	}

	data, err := s.client.Post(productBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return err
	}

	fmt.Println("---------", psResponse.Product)
	return nil
}
