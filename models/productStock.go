package models

type ProductStockData struct {
	IDProduct          string `json:"id_product,omitempty" xml:"id_product,omitempty"`
	IDProductAttribute string `json:"id_product_attribute,omitempty" xml:"id_product_attribute,omitempty"`
	IDShop             string `json:"id_shop,omitempty" xml:"id_shop,omitempty"`
	IDShopGroup        string `json:"id_shop_group,omitempty" xml:"id_shop_group,omitempty"`
	Quantity           string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	DependsOnStock     string `json:"depends_on_stock,omitempty" xml:"depends_on_stock,omitempty"`
	OutOfStock         string `json:"out_of_stock,omitempty" xml:"out_of_stock,omitempty"`
	Location           string `json:"location,omitempty" xml:"location,omitempty"`
}

type ProductStockReq struct {
	ID string `xml:"id,omitempty"`
	ProductStockData
}

type ProductStock struct {
	ID CustomID `json:"id,omitempty"`
	ProductStockData
}
