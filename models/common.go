package models

type MetaData struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type MetaAssociations struct {
	Id                 int `json:"id"`
	IdFeatureValue     int `json:"id_feature_value,omitempty"`
	IdProductAttribute int `json:"id_product_attribute,omitempty"`
}
