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

	resources, err := ps.Resource.ListResources()
	if err != nil {
		fmt.Println("----done-----", err)
		return
	}
	fmt.Println("----resources-----", resources)

	fmt.Println("----done-----")
}
