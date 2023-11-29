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

type ProductVariantService struct {
	client *Client
}

func newProductVariantService(client *Client) ProductVariantService {
	return ProductVariantService{
		client: client,
	}
}

func (s *ProductVariantService) ListProductVariant(productId string, limit, page int) (*ProductCombinationList, error) {
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

	return &productCombinationList, nil
}

func (s *ProductVariantService) CreateProductVariant(productVariant models.ProductVariantReq) (*models.Combination, error) {
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

func (s *ProductVariantService) UpdateProductVariant(productVariant models.ProductVariantReq) (*models.Combination, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Combinations: &productVariant})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Put(productVariantBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Combination, nil
}
