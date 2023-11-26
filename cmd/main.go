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

	//getProducts(ps)
	createProducts(ps)
	//createProductImage(ps)

	fmt.Println("----done-----")
}

func createProducts(ps *prestashopApi.PrestaShop) {
	product := models.ProductRequest{
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

	psss := models.Prestashop{
		Product: product,
	}
	fmt.Printf("----data---%#v--\n", psss)

	if err := ps.Product.CreateProduct(product); err != nil {
		fmt.Println("----done-----", err)
		return
	}

	fmt.Println("----resources-----")
}

func createProductImage(ps *prestashopApi.PrestaShop) {
	productId := 19
	filePath := "assets/test.jpg"

	// Open the image file
	//file, err := os.Open(filePath)
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//// Create a buffer to store the request body
	//var requestBody bytes.Buffer
	//writer := multipart.NewWriter(&requestBody)
	//
	//// Create a form field for the image file
	//part, err := writer.CreateFormFile("image", filePath)
	//if err != nil {
	//	fmt.Println("Error creating form file:", err)
	//	return
	//}
	//
	//// Copy the file content to the form field
	//_, err = io.Copy(part, file)
	//if err != nil {
	//	fmt.Println("Error copying file to form:", err)
	//	return
	//}
	//fmt.Println("FormDataContentType----------", writer.FormDataContentType())
	//
	//// Close the multipart writer to finalize the request body
	//writer.Close()

	if err := ps.Image.CreateProductImage(productId, filePath); err != nil {
		fmt.Println("----done-----", err)
		return
	}

	//Create the HTTP request
	//request, err := http.NewRequest("POST", "http://presto.local/api/images/products/19", &requestBody)
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//
	//// Set the Content-Type header
	//request.Header.Set("Content-Type", writer.FormDataContentType())
	//request.Header.Set("Output-Format", "JSON")
	//request.Header.Set("Io-Format", "JSON")
	//request.Header.Add("Accept", "application/json")
	//request.Header.Add("User-Agent", "test")
	//request.Header.Add("Output-Format", "JSON")
	//request.Header.Add("Io-Format", "JSON")
	//request.Header.Add("Authorization", fmt.Sprintf("Basic %s", "UVhZNTFDNlRXOUNUUURMWjI3NFVCQk5ENlpGNzZENEU6"))
	//
	//// Perform the request
	//client := &http.Client{}
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Println("Error performing request:", err)
	//	return
	//}
	//defer response.Body.Close()
	//
	//// Print the response status code and body
	//fmt.Println("Response Status:", response.Status)
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(response.Body)
	//fmt.Println("Response Body:", buf.String())

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
