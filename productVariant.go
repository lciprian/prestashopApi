package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productVariantBasePath = "combinations"

type ProductVariationList struct {
	Limit int
	Page  int
	Data  []models.Variant `json:"combinations,omitempty"`
}

type ProductVariantService struct {
	client *Client
}

func newProductVariantService(client *Client) ProductVariantService {
	return ProductVariantService{
		client: client,
	}
}

func (s *ProductVariantService) ListProductVariant(productId string, limit, page int) (*ProductVariationList, error) {
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

	productCombinationList := ProductVariationList{
		Limit: limit,
		Page:  page,
	}

	if err := s.client.Get(productVariantBasePath, queryParams, &productCombinationList); err != nil {
		return nil, err
	}

	return &productCombinationList, nil
}

func (s *ProductVariantService) CreateProductVariant(productVariant models.ProductVariantReq) (*models.Variant, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Combinations: &productVariant})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productVariantBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Variant, nil
}

func (s *ProductVariantService) UpdateProductVariant(productVariant models.ProductVariantReq) (*models.Variant, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{Combinations: &productVariant})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Put(productVariantBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Variant, nil
}
