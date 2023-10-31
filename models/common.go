package models

type MetaData struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type MetaAssociations struct {
	Id                 string `json:"id"`
	IdFeatureValue     string `json:"id_feature_value,omitempty"`
	IdProductAttribute string `json:"id_product_attribute,omitempty"`
}
