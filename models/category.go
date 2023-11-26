package models

type Category struct {
	Id                  int               `json:"id,omitempty"`
	IdParent            int               `json:"id_parent,omitempty"`
	LevelDepth          int               `json:"level_depth,omitempty"`
	NbProductsRecursive int               `json:"nb_products_recursive,omitempty"`
	Active              int               `json:"active,omitempty"`
	IdShopDefault       int               `json:"id_shop_default,omitempty"`
	IsRootCategory      int               `json:"is_root_category,omitempty"`
	Position            int               `json:"position,omitempty"`
	DateAdd             string            `json:"date_add,omitempty"`
	DateUpd             string            `json:"date_upd,omitempty"`
	Name                []MetaDataRequest `json:"name,omitempty"`
	LinkRewrite         []MetaDataRequest `json:"link_rewrite,omitempty"`
	Description         []MetaDataRequest `json:"description,omitempty"`
	MetaTitle           []MetaDataRequest `json:"meta_title,omitempty"`
	MetaDescription     []MetaDataRequest `json:"meta_description,omitempty"`
	MetaKeywords        []MetaDataRequest `json:"meta_keywords,omitempty"`
	Associations        struct {
		Categories [][]interface{} `json:"categories,omitempty"`
	}
}
