package main

import (
	"encoding/xml"
	"fmt"
	"github.com/lciprian/prestashopApi/models"

	"github.com/lciprian/prestashopApi"
)

func main() {
	appName := "Presta"
	urlStr := "https://presta.local"
	apiKey := "UVhZNTFDNlRXOUNUUURMWjI3NFVCQk5ENlpGNzZENEU6"

	ps := prestashopApi.NewPrestaShop(appName, urlStr, apiKey)
	fmt.Println("-----------", ps)

	//getResources(ps)

	//getProducts(ps)
	createProducts(ps)

	fmt.Println("----done-----")
}

func createProducts(ps *prestashopApi.PrestaShop) {
	product := models.Product2{
		New:   1,
		Price: "123",
		Name: []*models.MetaData{
			{
				models.Language{
					ID:   "1",
					Text: "My awesome Product",
				},
			},
		},
	}

	psss := models.Prestashop{
		Product: product,
	}
	fmt.Printf("----data---%#v--\n", psss)
	data, err := xml.Marshal(psss)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}
	fmt.Println("----data-----", string(data))
	//err := ps.Product.CreateProduct(product)
	//if err != nil {
	//	fmt.Println("----done-----", err)
	//	return
	//}

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

func getProducts(ps *prestashopApi.PrestaShop) {
	products, err := ps.Product.ListProducts(100, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", products)
}
