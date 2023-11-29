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
	ID           string       `json:"id"`
	IsColorGroup string       `json:"is_color_group"`
	GroupType    string       `json:"group_type"`
	Position     string       `json:"position"`
	Name         []Language   `json:"name>language"`
	PublicName   []Language   `json:"public_name>language"`
	Associations Associations `json:"associations"`
}

type Associations struct {
	ProductOptionValues []ProductOptionValue `json:"product_option_values"`
}
