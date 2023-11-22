package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/lciprian/prestashopApi"
	"github.com/lciprian/prestashopApi/models"
	"io"
	"mime/multipart"
	"os"
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
	createProductImage(ps)

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

func createProductImage(ps *prestashopApi.PrestaShop) {
	productId := 26
	filePath := "/home/cipi/Pictures/8kwallpapper.jpeg"

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form field for the image file
	part, err := writer.CreateFormFile("image", filePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file to form:", err)
		return
	}

	// Close the multipart writer to finalize the request body
	writer.Close()

	if err := ps.Image.CreateProductImage(productId, requestBody); err != nil {
		fmt.Println("----done-----", err)
		return
	}

	// Create the HTTP request
	//request, err := http.NewRequest("POST", "https://presta.local/api/images/products/25", &requestBody)
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//
	//// Set the Content-Type header
	//request.Header.Set("Content-Type", writer.FormDataContentType())
	//request.Header.Set("Output-Format", "JSON")
	//request.Header.Set("Io-Format", "JSON")
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
