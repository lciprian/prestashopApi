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

type ProductOptionValueList struct {
	Limit int
	Page  int
	Data  []models.ProductOptionValueL `json:"product_option_values,omitempty"`
}

type ProductOptionValueService struct {
	client *Client
}

func newProductOptionValueService(client *Client) ProductOptionValueService {
	return ProductOptionValueService{
		client: client,
	}
}

type ProductOptionValueCombinationList struct {
	Limit int
	Page  int
	Data  []models.Variant `json:"combinations,omitempty"`
}

func (s *ProductOptionValueService) ListProductOptionValues(prodOptionId string, limit, page int) (*ProductOptionValueList, error) {
	if page > 0 {
		page -= 1
	}
	offset := limit * page
	queryParams := url.Values{}
	queryParams.Add("display", "full")
	queryParams.Add("limit", fmt.Sprintf("%d,%d", offset, limit))

	if prodOptionId != "" {
		queryParams.Add("filter[id_attribute_group]", prodOptionId)
	}

	productOptionValueList := ProductOptionValueList{
		Limit: limit,
		Page:  page,
	}
	if err := s.client.Get(productOptionValueBasePath, queryParams, &productOptionValueList); err != nil {
		return nil, err
	}
	//	fmt.Println("-ListProducts---------", products)
	return &productOptionValueList, nil
}

func (s *ProductOptionValueService) CreateProductOptionValue(productOption models.ProductOptionValueReq) (*models.ProductOptionValue, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductOptionValue: &productOption})
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

	return &psResponse.ProductOptionValue, nil
}

func (s *ProductOptionValueService) UpdateProductOptionValue(productOption models.ProductOptionValueReq) (*models.ProductOptionValue, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductOptionValue: &productOption})
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

	return &psResponse.ProductOptionValue, nil
}
