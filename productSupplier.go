package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productProductSupplierBasePath = "product_suppliers"

type ProductSupplierList struct {
	Limit int
	Page  int
	Data  []models.ProductSupplierL `json:"product_suppliers,omitempty"`
}

type ProductSupplierService struct {
	client *Client
}

func newProductSupplierService(client *Client) ProductSupplierService {
	return ProductSupplierService{
		client: client,
	}
}

func (s *ProductSupplierService) GetProductSupplier(prodStockId string) (models.ProductSupplier, error) {
	queryParams := url.Values{}

	psResponse := Prestashop{}
	paPath := fmt.Sprintf("%s/%s", productProductSupplierBasePath, prodStockId)
	err := s.client.Get(paPath, queryParams, &psResponse)
	if err != nil {
		return models.ProductSupplier{}, err
	}

	//	fmt.Println("-ListProducts---------", products)
	return psResponse.ProductSupplier, nil
}

func (s *ProductSupplierService) CreateProductSupplier(productSupplier models.ProductSupplierReq) (*models.ProductSupplier, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductSupplier: &productSupplier})
	if err != nil {
		return nil, err
	}

	data, err := s.client.Post(productProductSupplierBasePath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return nil, err
	}

	return &psResponse.ProductSupplier, nil
}
