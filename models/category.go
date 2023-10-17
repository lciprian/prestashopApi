package models

type Category struct {
	Id                  int    `json:"id"`
	IdParent            int    `json:"id_parent"`
	LevelDepth          int    `json:"level_depth"`
	NbProductsRecursive int    `json:"nb_products_recursive"`
	Active              int    `json:"active"`
	IdShopDefault       int    `json:"id_shop_default"`
	IsRootCategory      int    `json:"is_root_category"`
	Position            int    `json:"position"`
	DateAdd             string `json:"date_add"`
	DateUpd             string `json:"date_upd"`
	Name                []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"name"`
	LinkRewrite []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"link_rewrite"`
	Description []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"description"`
	MetaTitle []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"meta_title"`
	MetaDescription []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"meta_description"`
	MetaKeywords []struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	} `json:"meta_keywords"`
	Associations struct {
		Categories [][]interface{} `json:"categories"`
	} `json:"associations"`
}
