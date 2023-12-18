package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

var productStockAvailableBasePath = "stock_availables"

type ProductStockAvailableList struct {
	Limit int
	Page  int
	Data  []models.ProductStockL `json:"product_option_values,omitempty"`
}

type ProductStockAvailableService struct {
	client *Client
}

func newProductStockAvailableService(client *Client) ProductStockAvailableService {
	return ProductStockAvailableService{
		client: client,
	}
}

func (s *ProductStockAvailableService) GetProductStockAvailable(prodStockId string) (models.ProductStock, error) {
	queryParams := url.Values{}

	psResponse := Prestashop{}
	paPath := fmt.Sprintf("%s/%s", productStockAvailableBasePath, prodStockId)
	err := s.client.Get(paPath, queryParams, &psResponse)
	if err != nil {
		return models.ProductStock{}, err
	}

	//	fmt.Println("-ListProducts---------", products)
	return psResponse.ProductStock, nil
}

func (s *ProductStockAvailableService) UpdateProductStockAvailable(productStock models.ProductStockReq) (models.ProductStock, error) {
	queryParams := url.Values{}

	buf, err := xml.Marshal(PrestashopReq{ProductStock: &productStock})
	if err != nil {
		return models.ProductStock{}, err
	}

	paPath := fmt.Sprintf("%s/%s", productStockAvailableBasePath, productStock.ID)
	data, err := s.client.Put(paPath, queryParams, bytes.NewBuffer(buf))
	if err != nil {
		return models.ProductStock{}, err
	}

	psResponse := Prestashop{}
	if err := json.Unmarshal(data, &psResponse); err != nil {
		return models.ProductStock{}, err
	}

	return psResponse.ProductStock, nil
}
