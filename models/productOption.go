package models

import (
	"encoding/xml"
)

type ProductOptionReq struct {
	XMLName      xml.Name   `xml:"product_option"`
	ID           string     `xml:"id"`
	IsColorGroup string     `xml:"is_color_group"`
	GroupType    string     `xml:"group_type"`
	Position     string     `xml:"position"`
	Name         []Language `xml:"name>language"`
	PublicName   []Language `xml:"public_name>language"`
}

type ProductOption struct {
	XMLName      xml.Name     `json:"product_option"`
	ID           int          `json:"id"`
	IsColorGroup string       `json:"is_color_group"`
	GroupType    string       `json:"group_type"`
	Position     string       `json:"position"`
	Name         []Language   `json:"name>language"`
	PublicName   []Language   `json:"public_name>language"`
	Associations Associations `json:"associations"`
}

type Associations struct {
	ProductOptionValues []ProductOptionValueAssociations `json:"product_option_values"`
}

type ProductOptionValueAssociations struct {
	ID               string     `json:"id,omitempty"`
	IDAttributeGroup string     `json:"id_attribute_group,omitempty"`
	Color            string     `json:"color,omitempty"`
	Position         string     `json:"position,omitempty"`
	Name             []Language `json:"name>language,omitempty"`
}
