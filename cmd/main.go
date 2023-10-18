package main

import (
	"fmt"

	"github.com/lciprian/prestashopApi"
)

func main() {
	appName := "Presta"
	urlStr := "https://presta.local"
	apiKey := "RUU4RVk5QUtGN0xORDZaN1FBVzVMTUQ1Q0czQkNUOUg6"

	ps := prestashopApi.NewPrestaShop(appName, urlStr, apiKey)
	fmt.Println("-----------", ps)

	//	getResources(ps)

	getProducts(ps)

	fmt.Println("----done-----")
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
	products, err := ps.Product.ListProducts(2, 1)
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Printf("----resources-----%#v", products[0])
}
