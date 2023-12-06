package models

type ProductOptionReq struct {
	ID           string        `xml:"id"`
	IsColorGroup string        `xml:"is_color_group"`
	GroupType    string        `xml:"group_type"`
	Position     string        `xml:"position"`
	Name         []LanguageReq `xml:"name>language"`
	PublicName   []LanguageReq `xml:"public_name>language"`
	Associations Associations  `xml:"associations"`
}

type ProductOption struct {
	ID           int           `json:"id"`
	IsColorGroup string        `json:"is_color_group"`
	GroupType    string        `json:"group_type"`
	Position     string        `json:"position"`
	Name         []LanguageReq `json:"name>language"`
	PublicName   []LanguageReq `json:"public_name>language"`
	Associations Associations  `json:"associations"`
}

type Associations struct {
	ProductOptionValues []ProductOptionValueAssociations `json:"product_option_values"`
}

type ProductOptionValueAssociations struct {
	ID               string        `json:"id,omitempty"`
	IDAttributeGroup string        `json:"id_attribute_group,omitempty"`
	Color            string        `json:"color,omitempty"`
	Position         string        `json:"position,omitempty"`
	Name             []LanguageReq `json:"name>language,omitempty"`
}
