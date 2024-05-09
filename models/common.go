package models

import (
	"encoding/xml"
)

type Prestashop struct {
	XMLName xml.Name   `xml:"prestashop"`
	Product ProductReq `xml:"product,omitempty"`
}

type MetaData struct {
	Id    string `json:"id" xml:"language,omitempty"`
	Value string `json:"value" xml:"description,omitempty"`
}

type MetaAssociations struct {
	ID                 string `json:"id" xml:"id,omitempty"`
	IdFeatureValue     string `json:"id_feature_value,omitempty" xml:"id_feature_value,omitempty"`
	IdProductAttribute string `json:"id_product_attribute,omitempty" xml:"id_product_attribute,omitempty"`
}

type MetaDataReq struct {
	Language []LanguageReq `xml:"language"`
}

type LanguageReq struct {
	ID   string `xml:"id,attr"`
	Text string `xml:",cdata"`
}

type Language struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type ResponseErrors struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
