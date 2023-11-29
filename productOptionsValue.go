package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productOptionValueBasePath = "product_option_values"

type ProductOptionValueService struct {
	client *Client
}

func newProductOptionValueService(client *Client) ProductOptionValueService {
	return ProductOptionValueService{
		client: client,
	}
}

type ProductOptionValueList struct {
	Limit int
	Page  int
	Data  []models.Product `json:"products,omitempty"`
}

type ProductOptionValueCombinationList struct {
	Limit int
	Page  int
	Data  []models.Combination `json:"combinations,omitempty"`
}

func newOptionValueService(client *Client) ProductService {
	return ProductService{
		client: client,
	}
}

func (s *ProductOptionValueService) ListProductOptionValues(prodOptionId string, limit, page int) (*ProductList, error) {
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

	if prodOptionId == "" {
		queryParams.Add("filter[id_attribute_group]", prodOptionId)
	}

	//products := make([]models.Product, 0)
	if err := s.client.Get(productOptionValueBasePath, queryParams, &productList); err != nil {
		return nil, err
	}
	//	fmt.Println("-ListProducts---------", products)
	return &productList, nil
}

func (s *ProductOptionValueService) CreateProductOptionValue(product models.ProductReq) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productOptionValueBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}

func (s *ProductOptionValueService) UpdateProductOptionValue(product models.ProductReq) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productOptionValueBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}
