package models

import "encoding/xml"

type Prestashop struct {
	XMLName xml.Name `xml:"prestashop"`
	Product Product2 `xml:"product,omitempty"`
}

//type Prestashop struct {
//	XMLName  xml.Name `xml:"prestashop"`
//	Category Category `xml:"category,omitempty"`
//	Product  Product  `xml:"product,omitempty"`
//}

//type MetaData struct {
//	Id    string `json:"id" xml:"language,omitempty"`
//	Value string `json:"value" xml:"description,omitempty"`
//}

type MetaAssociations struct {
	Id                 string `json:"id" `
	IdFeatureValue     string `json:"id_feature_value,omitempty"`
	IdProductAttribute string `json:"id_product_attribute,omitempty"`
}

type MetaData struct {
	Language Language `xml:"language"`
}

type Language struct {
	ID   string `xml:"id,attr"`
	Text string `xml:",cdata"`
}
