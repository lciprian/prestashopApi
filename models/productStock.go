package models

import (
	"encoding/json"
	"fmt"

	"github.com/lciprian/prestashopApi/helpers"
)

type ProductStock struct {
	ID                 string `json:"id,omitempty" xml:"id,omitempty"`
	IDProduct          string `json:"id_product,omitempty" xml:"id_product,omitempty"`
	IDProductAttribute string `json:"id_product_attribute,omitempty" xml:"id_product_attribute,omitempty"`
	IDShop             string `json:"id_shop,omitempty" xml:"id_shop,omitempty"`
	IDShopGroup        string `json:"id_shop_group,omitempty" xml:"id_shop_group,omitempty"`
	Quantity           string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	DependsOnStock     string `json:"depends_on_stock,omitempty" xml:"depends_on_stock,omitempty"`
	OutOfStock         string `json:"out_of_stock,omitempty" xml:"out_of_stock,omitempty"`
	Location           string `json:"location,omitempty" xml:"location,omitempty"`
}

func (m *ProductStock) UnmarshalJSON(data []byte) error {
	type productStock ProductStock

	type productStockRes struct {
		IdCustomer interface{} `json:"id,omitempty"`
		productStock
	}

	mReq := &productStockRes{}
	err := json.Unmarshal(data, &mReq)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if mReq.ID, err = helpers.IDtoString(mReq.IdCustomer); err != nil {
		return err
	}

	*m = ProductStock(mReq.productStock)

	return nil
}
