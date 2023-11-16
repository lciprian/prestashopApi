package models

import "encoding/json"

type Product struct {
	Id                      int                 `json:"id"`
	IdManufacturer          json.Number         `json:"id_manufacturer"`
	IdSupplier              json.Number         `json:"id_supplier"`
	IdCategoryDefault       json.Number         `json:"id_category_default"`
	New                     interface{}         `json:"new"`
	CacheDefaultAttribute   json.Number         `json:"cache_default_attribute"`
	IdDefaultImage          json.Number         `json:"id_default_image"`
	IdDefaultCombination    json.Number         `json:"id_default_combination"`
	IdTaxRulesGroup         json.Number         `json:"id_tax_rules_group"`
	PositionInCategory      json.Number         `json:"position_in_category"`
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
	DeliveryInStock         []MetaData          `json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        []MetaData          `json:"delivery_out_stock,omitempty"`
	ProductType             string              `json:"product_type"`
	OnSale                  string              `json:"on_sale"`
	OnlineOnly              string              `json:"online_only"`
	Ecotax                  string              `json:"ecotax"`
	MinimalQuantity         string              `json:"minimal_quantity"`
	LowStockThreshold       interface{}         `json:"low_stock_threshold"`
	LowStockAlert           string              `json:"low_stock_alert"`
	Price                   string              `json:"price"`
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
	MetaDescription         []MetaData          `json:"meta_description,omitempty"`
	MetaKeywords            []MetaData          `json:"meta_keywords,omitempty"`
	MetaTitle               []MetaData          `json:"meta_title,omitempty"`
	LinkRewrite             []MetaData          `json:"link_rewrite"`
	Name                    []MetaData          `json:"name,omitempty"`
	Description             []MetaData          `json:"description,omitempty"`
	DescriptionShort        []MetaData          `json:"description_short,omitempty"`
	AvailableNow            []MetaData          `json:"available_now,omitempty"`
	AvailableLater          []MetaData          `json:"available_later,omitempty"`
	Associations            ProductAssociations `json:"associations"`
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
	Id       int        `json:"id"`
	Position int        `json:"position"`
	Name     []MetaData `json:"name"`
}

type ProductFeatureValues struct {
	Id        int        `json:"id"`
	IdFeature int        `json:"id_feature"`
	Custom    int        `json:"custom"`
	Value     []MetaData `json:"value"`
}

type CustomizationFields struct {
	Id        int        `json:"id"`
	IdProduct int        `json:"id_product"`
	Type      int        `json:"type"`
	Required  int        `json:"required"`
	IsModule  int        `json:"is_module"`
	IsDeleted int        `json:"is_deleted"`
	Name      []MetaData `json:"name"`
}
