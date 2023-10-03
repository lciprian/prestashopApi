package main

import (
	"fmt"

	"github.com/lciprian/prestashopApi"
)

func main() {
	appName := "Presta"
	urlStr := "http://example.com"
	apiKey := "UCCLLQ9N2ARSHWCXLT74KUKSSK34BFKX"

	ps := prestashopApi.NewPrestaShop(appName, urlStr, apiKey)
	fmt.Println("-----------", ps)
	fmt.Println("----done-----")
}
