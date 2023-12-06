package models

import "encoding/json"

type ProductReq struct {
	Id string `json:"id" xml:"id,omitempty"`
	//IdManufacturer        json.Number `json:"id_manufacturer" xml:"id_manufacturer,omitempty"`
	//IdSupplier            json.Number `json:"id_supplier" xml:"id_supplier,omitempty"`
	IdCategoryDefault string      `json:"id_category_default" xml:"id_category_default,omitempty"`
	New               interface{} `json:"new" xml:"new,omitempty"`
	IdDefaultImage    string      `json:"id_default_image" xml:"id_default_image,omitempty"`
	//IdDefaultCombination  json.Number `json:"id_default_combination" xml:"id_default_combination,omitempty"`
	//IdTaxRulesGroup       json.Number `json:"id_tax_rules_group" xml:"id_tax_rules_group,omitempty"`
	//PositionInCategory    json.Number `json:"position_in_category" xml:"position_in_category,omitempty"`
	//ManufacturerName      interface{} `json:"manufacturer_name" xml:"manufacturer_name,omitempty"`
	Quantity      string `json:"quantity" xml:"quantity,omitempty"`
	Type          string `json:"type" xml:"type,omitempty"`
	IdShopDefault string `json:"id_shop_default" xml:"id_shop_default,omitempty"`
	Reference     string `json:"reference" xml:"reference,omitempty"`
	//SupplierReference     string      `json:"supplier_reference" xml:"supplier_reference,omitempty"`
	//Location              string      `json:"location" xml:"location,omitempty"`

	Ean13 string `json:"ean13" xml:"ean13,omitempty"`
	Isbn  string `json:"isbn" xml:"isbn,omitempty"`
	Upc   string `json:"upc" xml:"upc,omitempty"`
	Mpn   string `json:"mpn" xml:"mpn,omitempty"`

	Name             *MetaDataReq `json:"name,omitempty" xml:"name,omitempty"`
	Description      *MetaDataReq `json:"description,omitempty" xml:"description,omitempty"`
	DescriptionShort *MetaDataReq `json:"description_short,omitempty" xml:"description_short,omitempty"`
	//AvailableNow     *MetaDataReq         `json:"available_now,omitempty" xml:"available_now,omitempty"`
	//AvailableLater   *MetaDataReq         `json:"available_later,omitempty" xml:"available_later,omitempty"`
	//Associations     ProductAssociations `json:"associations" xml:"associations,omitempty"`
	MetaDescription *MetaDataReq `json:"meta_description,omitempty" xml:"meta_description,omitempty"`
	MetaKeywords    *MetaDataReq `json:"meta_keywords,omitempty" xml:"meta_keywords,omitempty"`
	MetaTitle       *MetaDataReq `json:"meta_title,omitempty" xml:"meta_title,omitempty"`

	Width  string `json:"width" xml:"width,omitempty"`
	Height string `json:"height" xml:"height,omitempty"`
	Depth  string `json:"depth" xml:"depth,omitempty"`
	Weight string `json:"weight" xml:"weight,omitempty"`
	//QuantityDiscount string `json:"quantity_discount" xml:"quantity_discount,omitempty"`

	Active      string `json:"active" xml:"active,omitempty"`
	State       string `json:"state" xml:"state,omitempty"`
	ProductType string `json:"product_type" xml:"product_type,omitempty"`
	//AdditionalDeliveryTimes string      `json:"additional_delivery_times" xml:"additional_delivery_times,omitempty"`
	//DeliveryInStock         *MetaDataReq  `json:"delivery_in_stock,omitempty" xml:"delivery_in_stock,omitempty"`
	//DeliveryOutStock        *MetaDataReq  `json:"delivery_out_stock,omitempty" xml:"delivery_out_stock,omitempty"`
	//OnSale                  string      `json:"on_sale" xml:"on_sale,omitempty"`
	//OnlineOnly              string      `json:"online_only" xml:"online_only,omitempty"`
	//Ecotax                  string      `json:"ecotax" xml:"ecotax,omitempty"`
	//MinimalQuantity         string      `json:"minimal_quantity" xml:"minimal_quantity,omitempty"`
	//LowStockThreshold       interface{} `json:"low_stock_threshold" xml:"low_stock_threshold,omitempty"`
	//LowStockAlert           string      `json:"low_stock_alert" xml:"low_stock_alert,omitempty"`

	Price string `json:"price" xml:"price"`
	//ShowPrice               string     `json:"show_price" xml:"show_price,omitempty"`
	//WholesalePrice          string     `json:"wholesale_price" xml:"wholesale_price,omitempty"`
	//Unity                   string     `json:"unity" xml:"unity,omitempty"`
	//UnitPriceRatio          string     `json:"unit_price_ratio" xml:"unit_price_ratio,omitempty"`
	//AdditionalShippingCost  string     `json:"additional_shipping_cost" xml:"additional_shipping_cost,omitempty"`
	//Customizable            string     `json:"customizable" xml:"customizable,omitempty"`
	//TextFields              string     `json:"text_fields" xml:"text_fields,omitempty"`
	//UploadableFiles         string     `json:"uploadable_files" xml:"uploadable_files,omitempty"`
	//RedirectType            string     `json:"redirect_type" xml:"redirect_type,omitempty"`
	//IdTypeRedirected        string     `json:"id_type_redirected" xml:"id_type_redirected,omitempty"`
	//AvailableForOrder       string     `json:"available_for_order" xml:"available_for_order,omitempty"`
	//AvailableDate           string     `json:"available_date" xml:"available_date,omitempty"`
	//ShowCondition           string     `json:"show_condition" xml:"show_condition,omitempty"`
	//Condition               string     `json:"condition" xml:"condition,omitempty"`
	//Indexed                 string     `json:"indexed" xml:"indexed,omitempty"`
	//Visibility              string     `json:"visibility" xml:"visibility,omitempty"`
	//AdvancedStockManagement string     `json:"advanced_stock_management" xml:"advanced_stock_management,omitempty"`
	//DateAdd                 string     `json:"date_add" xml:"date_add,omitempty"`
	//DateUpd                 string     `json:"date_upd" xml:"date_upd,omitempty"`
	//PackStockType           string     `json:"pack_stock_type" xml:"pack_stock_type,omitempty"`
	//LinkRewrite             []MetaDataReq `json:"link_rewrite" xml:"link_rewrite,omitempty"`
	//CacheIsPack             string     `json:"cache_is_pack" xml:"cache_is_pack,omitempty"`
	//CacheHasAttachments     string     `json:"cache_has_attachments" xml:"cache_has_attachments,omitempty"`
	//IsVirtual               string     `json:"is_virtual" xml:"is_virtual,omitempty"`
	//CacheDefaultAttribute json.Number `json:"cache_default_attribute" xml:"cache_default_attribute,omitempty"`
}

type ProductL struct {
	Id                      int                 `json:"id"`
	IdManufacturer          string              `json:"id_manufacturer"`
	IdSupplier              string              `json:"id_supplier"`
	IdCategoryDefault       string              `json:"id_category_default"`
	New                     interface{}         `json:"new"`
	CacheDefaultAttribute   string              `json:"cache_default_attribute"`
	IdDefaultImage          string              `json:"id_default_image"`
	IdDefaultCombination    json.Number         `json:"id_default_combination"`
	IdTaxRulesGroup         string              `json:"id_tax_rules_group"`
	PositionInCategory      string              `json:"position_in_category"`
	ManufacturerName        interface{}         `json:"manufacturer_name"`
	Quantity                string              `json:"quantity"`
	Type                    string              `json:"type"`
	IdShopDefault           json.Number         `json:"id_shop_default"`
	Reference               string              `json:"reference"`
	SupplierReference       string              `json:"supplier_reference"`
	Location                string              `json:"location"`
	Width                   string              `json:"width"`
	Height                  string              `json:"height"`
	Depth                   string              `json:"depth"`
	Weight                  string              `json:"weight"`
	QuantityDiscount        string              `json:"quantity_discount"`
	Ean13                   string              `json:"ean13"`
	Isbn                    string              `json:"isbn"`
	Upc                     string              `json:"upc"`
	Mpn                     string              `json:"mpn"`
	CacheIsPack             string              `json:"cache_is_pack"`
	CacheHasAttachments     string              `json:"cache_has_attachments"`
	IsVirtual               string              `json:"is_virtual"`
	State                   string              `json:"state"`
	AdditionalDeliveryTimes string              `json:"additional_delivery_times"`
	DeliveryInStock         []*MetaData         `json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        []*MetaData         `json:"delivery_out_stock,omitempty"`
	ProductType             string              `json:"product_type"`
	OnSale                  string              `json:"on_sale"`
	OnlineOnly              string              `json:"online_only"`
	Ecotax                  string              `json:"ecotax"`
	MinimalQuantity         string              `json:"minimal_quantity"`
	LowStockThreshold       interface{}         `json:"low_stock_threshold"`
	LowStockAlert           string              `json:"low_stock_alert"`
	Price                   string              `json:"price" xml:"price"`
	WholesalePrice          string              `json:"wholesale_price"`
	Unity                   string              `json:"unity"`
	UnitPriceRatio          string              `json:"unit_price_ratio"`
	AdditionalShippingCost  string              `json:"additional_shipping_cost"`
	Customizable            string              `json:"customizable"`
	TextFields              string              `json:"text_fields"`
	UploadableFiles         string              `json:"uploadable_files"`
	Active                  string              `json:"active"`
	RedirectType            string              `json:"redirect_type"`
	IdTypeRedirected        string              `json:"id_type_redirected"`
	AvailableForOrder       string              `json:"available_for_order"`
	AvailableDate           string              `json:"available_date"`
	ShowCondition           string              `json:"show_condition"`
	Condition               string              `json:"condition"`
	ShowPrice               string              `json:"show_price"`
	Indexed                 string              `json:"indexed"`
	Visibility              string              `json:"visibility"`
	AdvancedStockManagement string              `json:"advanced_stock_management"`
	DateAdd                 string              `json:"date_add"`
	DateUpd                 string              `json:"date_upd"`
	PackStockType           string              `json:"pack_stock_type"`
	MetaDescription         []*MetaData         `json:"meta_description,omitempty"`
	MetaKeywords            []*MetaData         `json:"meta_keywords,omitempty"`
	MetaTitle               []*MetaData         `json:"meta_title,omitempty"`
	LinkRewrite             []*MetaData         `json:"link_rewrite"`
	Name                    []*MetaData         `json:"name,omitempty"`
	Description             []*MetaData         `json:"description,omitempty"`
	DescriptionShort        []*MetaData         `json:"description_short,omitempty"`
	AvailableNow            []*MetaDataReq      `json:"available_now,omitempty"`
	AvailableLater          []*MetaDataReq      `json:"available_later,omitempty"`
	Associations            ProductAssociations `json:"associations"`
}

type Product struct {
	Id                      string      `json:"id"`
	IdManufacturer          string      `json:"id_manufacturer"`
	IdSupplier              string      `json:"id_supplier"`
	IdCategoryDefault       string      `json:"id_category_default"`
	New                     interface{} `json:"new"`
	CacheDefaultAttribute   string      `json:"cache_default_attribute"`
	IdDefaultImage          string      `json:"id_default_image"`
	IdDefaultCombination    json.Number `json:"id_default_combination"`
	IdTaxRulesGroup         string      `json:"id_tax_rules_group"`
	PositionInCategory      string      `json:"position_in_category"`
	ManufacturerName        interface{} `json:"manufacturer_name"`
	Quantity                string      `json:"quantity"`
	Type                    string      `json:"type"`
	IdShopDefault           json.Number `json:"id_shop_default"`
	Reference               string      `json:"reference"`
	SupplierReference       string      `json:"supplier_reference"`
	Location                string      `json:"location"`
	Width                   string      `json:"width"`
	Height                  string      `json:"height"`
	Depth                   string      `json:"depth"`
	Weight                  string      `json:"weight"`
	QuantityDiscount        string      `json:"quantity_discount"`
	Ean13                   string      `json:"ean13"`
	Isbn                    string      `json:"isbn"`
	Upc                     string      `json:"upc"`
	Mpn                     string      `json:"mpn"`
	CacheIsPack             string      `json:"cache_is_pack"`
	CacheHasAttachments     string      `json:"cache_has_attachments"`
	IsVirtual               string      `json:"is_virtual"`
	State                   string      `json:"state"`
	AdditionalDeliveryTimes string      `json:"additional_delivery_times"`
	DeliveryInStock         string      `json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        string      `json:"delivery_out_stock,omitempty"`
	ProductType             string      `json:"product_type"`
	OnSale                  string      `json:"on_sale"`
	OnlineOnly              string      `json:"online_only"`
	Ecotax                  string      `json:"ecotax"`
	MinimalQuantity         string      `json:"minimal_quantity"`
	LowStockThreshold       interface{} `json:"low_stock_threshold"`
	LowStockAlert           string      `json:"low_stock_alert"`
	Price                   string      `json:"price" xml:"price"`
	WholesalePrice          string      `json:"wholesale_price"`
	Unity                   string      `json:"unity"`
	UnitPriceRatio          string      `json:"unit_price_ratio"`
	AdditionalShippingCost  string      `json:"additional_shipping_cost"`
	Customizable            string      `json:"customizable"`
	TextFields              string      `json:"text_fields"`
	UploadableFiles         string      `json:"uploadable_files"`
	Active                  string      `json:"active"`
	RedirectType            string      `json:"redirect_type"`
	IdTypeRedirected        string      `json:"id_type_redirected"`
	AvailableForOrder       string      `json:"available_for_order"`
	AvailableDate           string      `json:"available_date"`
	ShowCondition           string      `json:"show_condition"`
	Condition               string      `json:"condition"`
	ShowPrice               string      `json:"show_price"`
	Indexed                 string      `json:"indexed"`
	Visibility              string      `json:"visibility"`
	AdvancedStockManagement string      `json:"advanced_stock_management"`
	DateAdd                 string      `json:"date_add"`
	DateUpd                 string      `json:"date_upd"`
	PackStockType           string      `json:"pack_stock_type"`
	MetaDescription         string      `json:"meta_description,omitempty"`
	MetaKeywords            string      `json:"meta_keywords,omitempty"`
	MetaTitle               string      `json:"meta_title,omitempty"`
	//LinkRewrite             []*MetaData         `json:"link_rewrite,omitempty"`
	Name             []*MetaData         `json:"name,omitempty"`
	Description      []*MetaData         `json:"description,omitempty"`
	DescriptionShort []*MetaData         `json:"description_short,omitempty"`
	AvailableNow     string              `json:"available_now,omitempty"`
	AvailableLater   string              `json:"available_later,omitempty"`
	Associations     ProductAssociations `json:"associations"`
}

type ProductAssociations struct {
	Categories          []MetaAssociations `json:"categories,omitempty"`
	Images              []MetaAssociations `json:"images,omitempty"`
	Combinations        []MetaAssociations `json:"combinations,omitempty"`
	ProductOptionValues []MetaAssociations `json:"product_option_values,omitempty"`
	ProductFeatures     []MetaAssociations `json:"product_features,omitempty"`
	StockAvailables     []MetaAssociations `json:"stock_availables,omitempty"`
}

type ProductFeatures struct {
	Id       int           `json:"id"`
	Position int           `json:"position"`
	Name     []MetaDataReq `json:"name"`
}

type ProductFeatureValues struct {
	Id        int           `json:"id"`
	IdFeature int           `json:"id_feature"`
	Custom    int           `json:"custom"`
	Value     []MetaDataReq `json:"value"`
}

type CustomizationFields struct {
	Id        int           `json:"id"`
	IdProduct int           `json:"id_product"`
	Type      int           `json:"type"`
	Required  int           `json:"required"`
	IsModule  int           `json:"is_module"`
	IsDeleted int           `json:"is_deleted"`
	Name      []MetaDataReq `json:"name"`
}
