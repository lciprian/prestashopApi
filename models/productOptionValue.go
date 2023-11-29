package models

import (
	"encoding/xml"
)

type ProductOptionValueReq struct {
	XMLName          xml.Name   `xml:"product_option_value"`
	ID               string     `xml:"id"`
	IDAttributeGroup string     `xml:"id_attribute_group"`
	Color            string     `xml:"color"`
	Position         string     `xml:"position"`
	Name             []Language `xml:"name>language"`
}

type ProductOptionValue struct {
	XMLName          xml.Name   `json:"product_option_value"`
	ID               string     `json:"id,omitempty"`
	IDAttributeGroup string     `json:"id_attribute_group,omitempty"`
	Color            string     `json:"color,omitempty"`
	Position         string     `json:"position,omitempty"`
	Name             []Language `json:"name>language,omitempty"`
}
