package models

type ProductSupplierReq struct {
	ID                       string `xml:"id,omitempty"`
	IDProduct                string `xml:"id_product,omitempty"`
	IDProductAttribute       string `xml:"id_product_attribute,omitempty"`
	IDSupplier               string `xml:"id_supplier,omitempty"`
	IDCurrency               string `xml:"id_currency,omitempty"`
	ProductSupplierReference string `xml:"product_supplier_reference,omitempty"`
	ProductSupplierPriceTe   string `xml:"product_supplier_price_te,omitempty"`
}

type ProductSupplierL struct {
	ID                       int    `json:"id,omitempty"`
	IDProduct                string `json:"id_product,omitempty"`
	IDProductAttribute       string `json:"id_product_attribute,omitempty"`
	IDSupplier               string `json:"id_supplier,omitempty"`
	IDCurrency               string `json:"id_currency,omitempty"`
	ProductSupplierReference string `json:"product_supplier_reference,omitempty"`
	ProductSupplierPriceTe   string `json:"product_supplier_price_te,omitempty"`
}

type ProductSupplier struct {
	ID                       string `json:"id,omitempty"`
	IDProduct                string `json:"id_product,omitempty"`
	IDProductAttribute       string `json:"id_product_attribute,omitempty"`
	IDSupplier               string `json:"id_supplier,omitempty"`
	IDCurrency               string `json:"id_currency,omitempty"`
	ProductSupplierReference string `json:"product_supplier_reference,omitempty"`
	ProductSupplierPriceTe   string `json:"product_supplier_price_te,omitempty"`
}
