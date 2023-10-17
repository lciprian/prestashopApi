package models

type Product struct {
	Id                      int                 `json:"id"`
	IdManufacturer          int                 `json:"id_manufacturer"`
	IdSupplier              int                 `json:"id_supplier"`
	IdCategoryDefault       int                 `json:"id_category_default"`
	New                     interface{}         `json:"new"`
	CacheDefaultAttribute   int                 `json:"cache_default_attribute"`
	IdDefaultImage          int                 `json:"id_default_image"`
	IdDefaultCombination    int                 `json:"id_default_combination"`
	IdTaxRulesGroup         int                 `json:"id_tax_rules_group"`
	PositionInCategory      int                 `json:"position_in_category"`
	ManufacturerName        string              `json:"manufacturer_name"`
	Quantity                int                 `json:"quantity"`
	Type                    string              `json:"type"`
	IdShopDefault           int                 `json:"id_shop_default"`
	Reference               string              `json:"reference"`
	SupplierReference       string              `json:"supplier_reference"`
	Location                string              `json:"location"`
	Width                   string              `json:"width"`
	Height                  string              `json:"height"`
	Depth                   string              `json:"depth"`
	Weight                  string              `json:"weight"`
	QuantityDiscount        int                 `json:"quantity_discount"`
	Ean13                   string              `json:"ean13"`
	Isbn                    string              `json:"isbn"`
	Upc                     string              `json:"upc"`
	Mpn                     string              `json:"mpn"`
	CacheIsPack             int                 `json:"cache_is_pack"`
	CacheHasAttachments     int                 `json:"cache_has_attachments"`
	IsVirtual               int                 `json:"is_virtual"`
	State                   int                 `json:"state"`
	AdditionalDeliveryTimes int                 `json:"additional_delivery_times"`
	DeliveryInStock         []MetaData          `json:"delivery_in_stock,omitempty"`
	DeliveryOutStock        []MetaData          `json:"delivery_out_stock,omitempty"`
	ProductType             string              `json:"product_type"`
	OnSale                  int                 `json:"on_sale"`
	OnlineOnly              int                 `json:"online_only"`
	Ecotax                  string              `json:"ecotax"`
	MinimalQuantity         int                 `json:"minimal_quantity"`
	LowStockThreshold       interface{}         `json:"low_stock_threshold"`
	LowStockAlert           int                 `json:"low_stock_alert"`
	Price                   string              `json:"price"`
	WholesalePrice          string              `json:"wholesale_price"`
	Unity                   string              `json:"unity"`
	UnitPriceRatio          string              `json:"unit_price_ratio"`
	AdditionalShippingCost  string              `json:"additional_shipping_cost"`
	Customizable            int                 `json:"customizable"`
	TextFields              int                 `json:"text_fields"`
	UploadableFiles         int                 `json:"uploadable_files"`
	Active                  int                 `json:"active"`
	RedirectType            string              `json:"redirect_type"`
	IdTypeRedirected        int                 `json:"id_type_redirected"`
	AvailableForOrder       int                 `json:"available_for_order"`
	AvailableDate           string              `json:"available_date"`
	ShowCondition           int                 `json:"show_condition"`
	Condition               string              `json:"condition"`
	ShowPrice               int                 `json:"show_price"`
	Indexed                 int                 `json:"indexed"`
	Visibility              string              `json:"visibility"`
	AdvancedStockManagement int                 `json:"advanced_stock_management"`
	DateAdd                 string              `json:"date_add"`
	DateUpd                 string              `json:"date_upd"`
	PackStockType           int                 `json:"pack_stock_type"`
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
