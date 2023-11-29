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

type ProductOptionService struct {
	client *Client
}

func newProductOptionService(client *Client) ProductOptionService {
	return ProductOptionService{
		client: client,
	}
}

type ProductOptionList struct {
	Limit int
	Page  int
	Data  []models.Product `json:"products,omitempty"`
}

func (s *ProductService) ListProductOptions(limit, page int) (*ProductList, error) {
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
	if err := s.client.Get(productOptionBasePath, queryParams, &productList); err != nil {
		return nil, err
	}
	//	fmt.Println("-ListProducts---------", products)
	return &productList, nil
}

func (s *ProductService) CreateProductOption(product models.ProductRequest) (*models.Product, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(Prestashop{Product: &product})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productOptionBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop2{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.Product, nil
}