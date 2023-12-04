package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productOptionBasePath = "product_options"

type ProductOptionList struct {
	Limit int
	Page  int
	Data  []models.ProductOption `json:"product_options,omitempty"`
}

type ProductOptionService struct {
	client *Client
}

func newProductOptionService(client *Client) ProductOptionService {
	return ProductOptionService{
		client: client,
	}
}

func (s *ProductOptionService) ListProductOptions(prodOptionId string, limit, page int) (*ProductOptionList, error) {
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

	result := ProductOptionList{
		Limit: limit,
		Page:  page,
	}

	if err := s.client.Get(productOptionBasePath, queryParams, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ProductOptionService) CreateProductOption(productOption models.ProductOptionReq) (*models.ProductOption, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductOption: &productOption})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productOptionBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.ProductOption, nil
}

func (s *ProductOptionService) UpdateProductOption(productOption models.ProductOptionReq) (*models.ProductOption, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductOption: &productOption})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productOptionBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.ProductOption, nil
}
