package models

type ProductOptionValueReq struct {
	ID               string        `xml:"id"`
	IDAttributeGroup string        `xml:"id_attribute_group"`
	Color            string        `xml:"color"`
	Position         string        `xml:"position"`
	Name             []LanguageReq `xml:"name>language"`
}

type ProductOptionValueL struct {
	ID               int    `json:"id,omitempty"`
	IDAttributeGroup string `json:"id_attribute_group,omitempty"`
	Color            string `json:"color,omitempty"`
	Position         string `json:"position,omitempty"`
	Name             string `json:"name,omitempty"`
}

type ProductOptionValue struct {
	ID               string `json:"id,omitempty"`
	IDAttributeGroup string `json:"id_attribute_group,omitempty"`
	Color            string `json:"color,omitempty"`
	Position         int    `json:"position,omitempty"`
	Name             string `json:"name,omitempty"`
}
