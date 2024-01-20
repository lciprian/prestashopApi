package main

import (
	"fmt"

	"github.com/lciprian/prestashopApi"
	"github.com/lciprian/prestashopApi/models"
)

func main() {
	appName := "Presta"
	urlStr := "http://presto.local"
	apiKey := "UVhZNTFDNlRXOUNUUURMWjI3NFVCQk5ENlpGNzZENEU6"

	ps := prestashopApi.NewPrestaShop(appName, urlStr, apiKey)
	fmt.Println("-----------", ps)

	//getResources(ps)

	//	getProducts(ps)
	//createProducts(ps)
	//getProduct(ps, "116")
	//updateProducts(ps, "45")
	//createProductImage(ps)

	getProductVariants(ps, "126")
	//createProductVariant(ps, "116")
	//updateProductVariant(ps, "45", "44")

	//getProductOptions(ps, "2")

	//getProductOptionValue(ps, "2")
	//createProductOptionValue(ps, "2")
	//updateProductOptionValue(ps, "2")

	//getProductStock(ps, "100")

	//updateProductStock(ps, "230", "80")

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

func getProduct(ps *prestashopApi.PrestaShop, pId string) {

	product, err := ps.Product.GetProduct(pId)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----getProduct-----\n", product)
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

func createProductVariant(ps *prestashopApi.PrestaShop, pId string) {
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

	productVariant, err := ps.ProductVariant.UpdateProductVariant(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func deleteProductVariant(ps *prestashopApi.PrestaShop, pvId string) {
	if err := ps.ProductVariant.DeleteProductVariant(pvId); err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----success-----")
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

func getProductOption(ps *prestashopApi.PrestaShop, pId string) {
	result, err := ps.ProductOptionValue.ListProductOptionValues(pId, 100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", result)
}

func createProductOption(ps *prestashopApi.PrestaShop, pId string) {
	productOptionReq := getNewProductOption()
	fmt.Printf("----data---%#v--\n", productOptionReq)

	productVariant, err := ps.ProductOption.CreateProductOption(productOptionReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func updateProductOption(ps *prestashopApi.PrestaShop, pId, vId string) {
	productOptionReq := getUpdateProductOption()
	fmt.Printf("----data---%#v--\n", productOptionReq)

	productOption, err := ps.ProductOption.UpdateProductOption(productOptionReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productOption)
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
	productOptionReq := getNewProductOptionValue(pId)
	fmt.Printf("----data---%#v--\n", productOptionReq)

	productVariant, err := ps.ProductOptionValue.CreateProductOptionValue(productOptionReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func updateProductOptionValue(ps *prestashopApi.PrestaShop, pId, vId string) {
	productVariantReq := getUpdateProductOptionValue(pId, vId)
	fmt.Printf("----data---%#v--\n", productVariantReq)

	productVariant, err := ps.ProductOptionValue.UpdateProductOptionValue(productVariantReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", productVariant)
}

func getProductStock(ps *prestashopApi.PrestaShop, prodStockId string) {
	result, err := ps.ProductStock.GetProductStockAvailable(prodStockId)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", result)
}

func updateProductStock(ps *prestashopApi.PrestaShop, prodStockId, productId string) {
	productStockReq := getUpdateProductStock(prodStockId, productId)
	fmt.Printf("----data---%#v--\n", productStockReq)

	prodStock, err := ps.ProductStock.UpdateProductStockAvailable(productStockReq)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----", prodStock)
}

func getNewProduct() models.ProductReq {
	return models.ProductReq{
		New:               1,
		IdShopDefault:     "2",
		Active:            "1",
		State:             "1",
		AvailableForOrder: "1",
		Type:              "0",
		Price:             "123",
		IdCategoryDefault: "9",
		Associations: models.ProductAssociations{
			Categories: []models.MetaAssociations{
				{
					ID: "2",
				},
				{
					ID: "9",
				},
			},
		},
		Name: &models.MetaDataReq{
			Language: []models.LanguageReq{
				{
					ID:   "1",
					Text: "My awesome Product 31",
				},
				{
					ID:   "2",
					Text: "My awesome Product 31",
				},
			},
		},
		Description: &models.MetaDataReq{
			Language: []models.LanguageReq{
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
		DescriptionShort: &models.MetaDataReq{
			Language: []models.LanguageReq{
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
		Quantity:        100,
		Reference:       fmt.Sprintf("testsku1-%s", productId),
		Price:           "123",
		Weight:          "100",
		ProductOptionValue: struct {
			ID string `xml:"id"`
		}{
			ID: "40",
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
			ID: "26",
		},
	}
}

func getNewProductOption() models.ProductOptionReq {
	return models.ProductOptionReq{
		Name: []models.LanguageReq{
			{
				ID:   "1",
				Text: "Marsala",
			},
			{
				ID:   "2",
				Text: "Marsala",
			},
		},
		PublicName: []models.LanguageReq{
			{
				ID:   "1",
				Text: "Marsala",
			},
			{
				ID:   "2",
				Text: "Marsala",
			},
		},
	}
}

func getUpdateProductOption() models.ProductOptionReq {
	return models.ProductOptionReq{
		Name: []models.LanguageReq{
			{
				ID:   "2",
				Text: "Marsala",
			},
		},
		PublicName: []models.LanguageReq{
			{
				ID:   "2",
				Text: "Marsala",
			},
		},
	}
}

func getNewProductOptionValue(attrId string) models.ProductOptionValueReq {
	return models.ProductOptionValueReq{
		IDAttributeGroup: attrId,
		Color:            "Marsala",
		Name: []models.LanguageReq{
			{
				ID:   "1",
				Text: "Marsala",
			},
			{
				ID:   "2",
				Text: "Marsala",
			},
		},
	}
}

func getUpdateProductOptionValue(variantId, productId string) models.ProductOptionValueReq {
	return models.ProductOptionValueReq{}
}

func getUpdateProductStock(prodStockId, productId string) models.ProductStockReq {
	return models.ProductStockReq{
		ID:                 prodStockId,
		IDProduct:          productId,
		IDProductAttribute: "1",
		IDShop:             "1",
		Quantity:           "10",
		DependsOnStock:     "0",
		OutOfStock:         "2",
	}
}
