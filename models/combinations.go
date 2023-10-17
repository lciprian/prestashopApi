package models

type Combinations struct {
	Id                int         `json:"id"`
	IdProduct         int         `json:"id_product"`
	Location          string      `json:"location"`
	Ean13             string      `json:"ean13"`
	Isbn              string      `json:"isbn"`
	Upc               string      `json:"upc"`
	Mpn               string      `json:"mpn"`
	Quantity          int         `json:"quantity"`
	Reference         string      `json:"reference"`
	SupplierReference string      `json:"supplier_reference"`
	WholesalePrice    string      `json:"wholesale_price"`
	Price             string      `json:"price"`
	Ecotax            string      `json:"ecotax"`
	Weight            string      `json:"weight"`
	UnitPriceImpact   string      `json:"unit_price_impact"`
	MinimalQuantity   int         `json:"minimal_quantity"`
	LowStockThreshold interface{} `json:"low_stock_threshold"`
	LowStockAlert     int         `json:"low_stock_alert"`
	DefaultOn         interface{} `json:"default_on"`
	AvailableDate     string      `json:"available_date"`
	Associations      struct {
		ProductOptionValues [][]interface{} `json:"product_option_values"`
		Images              [][]interface{} `json:"images"`
	} `json:"associations"`
}
