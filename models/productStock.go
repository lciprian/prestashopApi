package models

type ProductStockReq struct {
	ID                 string `xml:"id,omitempty"`
	IDProduct          string `xml:"id_product,omitempty"`
	IDProductAttribute string `xml:"id_product_attribute,omitempty"`
	IDShop             string `xml:"id_shop,omitempty"`
	IDShopGroup        string `xml:"id_shop_group,omitempty"`
	Quantity           string `xml:"quantity,omitempty"`
	DependsOnStock     string `xml:"depends_on_stock,omitempty"`
	OutOfStock         string `xml:"out_of_stock,omitempty"`
	Location           string `xml:"location,omitempty"`
}

type ProductStockL struct {
	ID                 int    `json:"id,omitempty"`
	IDProduct          string `json:"id_product,omitempty"`
	IDProductAttribute string `json:"id_product_attribute,omitempty"`
	IDShop             string `json:"id_shop,omitempty"`
	IDShopGroup        string `json:"id_shop_group,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	DependsOnStock     string `json:"depends_on_stock,omitempty"`
	OutOfStock         string `json:"out_of_stock,omitempty"`
	Location           string `json:"location,omitempty"`
}

type ProductStock struct {
	ID                 string `json:"id,omitempty"`
	IDProduct          string `json:"id_product,omitempty"`
	IDProductAttribute string `json:"id_product_attribute,omitempty"`
	IDShop             string `json:"id_shop,omitempty"`
	IDShopGroup        string `json:"id_shop_group,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	DependsOnStock     string `json:"depends_on_stock,omitempty"`
	OutOfStock         string `json:"out_of_stock,omitempty"`
	Location           string `json:"location,omitempty"`
}
