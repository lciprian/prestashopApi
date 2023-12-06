package models

type ProductVariantReq struct {
	Id                 string      `json:"id" xml:"id,omitempty"`
	IdProduct          string      `json:"id_product" xml:"id_product,omitempty"`
	Location           string      `json:"location" xml:"location,omitempty"`
	Ean13              string      `json:"ean13" xml:"ean13,omitempty"`
	Isbn               string      `json:"isbn" xml:"isbn,omitempty"`
	Upc                string      `json:"upc" xml:"upc,omitempty"`
	Mpn                string      `json:"mpn" xml:"mpn,omitempty"`
	Quantity           int         `json:"quantity" xml:"quantity,omitempty"`
	Reference          string      `json:"reference" xml:"reference,omitempty"`
	SupplierReference  string      `json:"supplier_reference" xml:"supplier_reference,omitempty"`
	WholesalePrice     string      `json:"wholesale_price" xml:"wholesale_price,omitempty"`
	Price              string      `json:"price" xml:"price,omitempty"`
	Ecotax             string      `json:"ecotax" xml:"ecotax,omitempty"`
	Weight             string      `json:"weight" xml:"weight,omitempty"`
	UnitPriceImpact    string      `json:"unit_price_impact" xml:"unit_price_impact,omitempty"`
	MinimalQuantity    int         `json:"minimal_quantity" xml:"minimal_quantity,omitempty"`
	LowStockThreshold  interface{} `json:"low_stock_threshold" xml:"low_stock_threshold,omitempty"`
	LowStockAlert      int         `json:"low_stock_alert" xml:"low_stock_alert,omitempty"`
	DefaultOn          interface{} `json:"default_on" xml:"default_on,omitempty"`
	AvailableDate      string      `json:"available_date" xml:"available_date,omitempty"`
	ProductOptionValue struct {
		ID string `xml:"id"`
	} `xml:"associations>product_option_values>product_option_value,omitempty"`
	Image struct {
		ID string `xml:"id"`
	} `xml:"associations>images>image,omitempty"`
}

type VariantL struct {
	Id                int         `json:"id"`
	IdProduct         string      `json:"id_product"`
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
	MinimalQuantity   string      `json:"minimal_quantity"`
	LowStockThreshold interface{} `json:"low_stock_threshold"`
	LowStockAlert     string      `json:"low_stock_alert"`
	DefaultOn         interface{} `json:"default_on"`
	AvailableDate     string      `json:"available_date"`
	Associations      struct {
		ProductOptionValues []struct {
			ID string `json:"id"`
		} `json:"product_option_values"`
		Images []struct {
			ID string `json:"id"`
		} `json:"images"`
	} `json:"associations"`
}

type Variant struct {
	Id                string      `json:"id"`
	IdProduct         string      `json:"id_product"`
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
	MinimalQuantity   string      `json:"minimal_quantity"`
	LowStockThreshold interface{} `json:"low_stock_threshold"`
	LowStockAlert     string      `json:"low_stock_alert"`
	DefaultOn         interface{} `json:"default_on"`
	AvailableDate     string      `json:"available_date"`
	Associations      struct {
		ProductOptionValues []struct {
			ID string `json:"id"`
		} `json:"product_option_values"`
		Images []struct {
			ID string `json:"id"`
		} `json:"images"`
	} `json:"associations"`
}
