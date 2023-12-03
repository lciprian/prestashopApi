package main

import (
	"fmt"
	"github.com/lciprian/prestashopApi/models"

	"github.com/lciprian/prestashopApi"
)

func main() {
	appName := "Presta"
	urlStr := "http://presto.local"
	apiKey := "UVhZNTFDNlRXOUNUUURMWjI3NFVCQk5ENlpGNzZENEU6"

	ps := prestashopApi.NewPrestaShop(appName, urlStr, apiKey)
	fmt.Println("-----------", ps)

	//getResources(ps)

	//getProducts(ps)
	//createProducts(ps)
	//updateProducts(ps, "45")
	//getProductVariants(ps, "44")
	//createProductVariant(ps, "44")
	//updateProductVariant(ps, "45", "44")
	//createProductImage(ps)
	getProductOptions(ps, "2")
	//getProductOptionValue(ps, "2")
	createProductOptionValue(ps, "2")

	fmt.Println("----done-----")
}

func getProducts(ps *prestashopApi.PrestaShop) {
	products, err := ps.Product.ListProducts(100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", products)
}

func createProducts(ps *prestashopApi.PrestaShop) {
	productReq := getNewProduct()

	//ProductType combinations
	fmt.Printf("----data---%#v--\n", productReq)

	product, err := ps.Product.CreateProduct(productReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----product-----\n", product)
}

func updateProducts(ps *prestashopApi.PrestaShop, pId string) {
	productReq := getNewProduct()
	productReq.Id = pId
	productReq.Reference = "REF112"
	//ProductType combinations
	fmt.Printf("----data---%#v--\n", productReq)

	product, err := ps.Product.UpdateProduct(productReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----product-----\n", product)
}

func getProductVariants(ps *prestashopApi.PrestaShop, pId string) {
	variants, err := ps.ProductVariant.ListProductVariant(pId, 100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", variants)
}

func createProductVariants(ps *prestashopApi.PrestaShop, pId string) {
	productVariantReq := getNewProductVariant(pId)
	fmt.Printf("----data---%#v--\n", productVariantReq)

	productVariant, err := ps.ProductVariant.CreateProductVariant(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func updateProductVariant(ps *prestashopApi.PrestaShop, pId, vId string) {
	productVariantReq := getNewUpdateProductVariant(vId, pId)
	fmt.Printf("----data---%#v--\n", productVariantReq)

	productVariant, err := ps.ProductVariant.CreateProductVariant(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func createProductImage(ps *prestashopApi.PrestaShop) {
	productId := 19
	filePath := "assets/test.jpg"

	if err := ps.Image.CreateProductImage(productId, filePath); err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----")
}

func getResources(ps *prestashopApi.PrestaShop) {
	resources, err := ps.Resource.ListResources()
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", resources)
}

func getProductOptions(ps *prestashopApi.PrestaShop, pId string) {
	result, err := ps.ProductOption.ListProductOptions(pId, 100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", result)
}

func getProductOptionValue(ps *prestashopApi.PrestaShop, pId string) {
	result, err := ps.ProductOptionValue.ListProductOptionValues(pId, 100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", result)
}

func createProductOptionValue(ps *prestashopApi.PrestaShop, pId string) {
	productVariantReq := getNewProductVariant(pId)
	fmt.Printf("----data---%#v--\n", productVariantReq)

	productVariant, err := ps.ProductOptionValue.CreateProductOptionValue(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func updateProductOptionValue(ps *prestashopApi.PrestaShop, pId, vId string) {
	productVariantReq := getNewUpdateProductVariant(vId, pId)
	fmt.Printf("----data---%#v--\n", productVariantReq)

	productVariant, err := ps.ProductOptionValue.CreateProductOptionValue(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func getNewProduct() models.ProductReq {
	return models.ProductReq{
		New:               1,
		IdCategoryDefault: "8",
		IdShopDefault:     "2",
		Active:            "1",
		State:             "1",
		Type:              "0",
		Price:             "123",
		Name: &models.MetaDataRequest{
			Language: []models.Language{
				{
					ID:   "1",
					Text: "My awesome Product 2",
				},
				{
					ID:   "2",
					Text: "My awesome Product 2",
				},
			},
		},
		Description: &models.MetaDataRequest{
			Language: []models.Language{
				{
					ID:   "1",
					Text: "My awesome Product Description 2",
				},
				{
					ID:   "2",
					Text: "My awesome Product Description 2",
				},
			},
		},
		DescriptionShort: &models.MetaDataRequest{
			Language: []models.Language{
				{
					ID:   "1",
					Text: "My awesome Product DescriptionShort 2",
				},
				{
					ID:   "2",
					Text: "My awesome Product DescriptionShort 2",
				},
			},
		},
		Weight: "100",
		Height: "150",
		Depth:  "50",
	}
}

func getNewProductVariant(productId string) models.ProductVariantReq {
	return models.ProductVariantReq{
		IdProduct:       productId,
		MinimalQuantity: 1,
		Reference:       "testsku1",
		Price:           "123",
		Weight:          "100",
		ProductOptionValue: struct {
			ID string `xml:"id"`
		}{
			ID: "5",
		},
	}
}

func getNewUpdateProductVariant(variantId, productId string) models.ProductVariantReq {
	return models.ProductVariantReq{
		Id:              variantId,
		IdProduct:       productId,
		MinimalQuantity: 1,
		Reference:       "testsku1",
		Price:           "123",
		Weight:          "100",
		ProductOptionValue: struct {
			ID string `xml:"id"`
		}{
			ID: "6",
		},
	}
}

func getNewProductOptionValue(productId string) models.ProductVariantReq {
	return models.ProductVariantReq{
		IdProduct:       productId,
		MinimalQuantity: 1,
		Reference:       "testsku1",
		Price:           "123",
		Weight:          "100",
		ProductOptionValue: struct {
			ID string `xml:"id"`
		}{
			ID: "5",
		},
	}
}

func getNewUpdateProductOptionValue(variantId, productId string) models.ProductVariantReq {
	return models.ProductVariantReq{
		Id:              variantId,
		IdProduct:       productId,
		MinimalQuantity: 1,
		Reference:       "testsku1",
		Price:           "123",
		Weight:          "100",
		ProductOptionValue: struct {
			ID string `xml:"id"`
		}{
			ID: "6",
		},
	}
}
