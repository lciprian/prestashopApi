package prestashopApi

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

var imagesProductBasePath = "images/products/%d"

type ImageService struct {
	client *Client
}

func newImageService(client *Client) ImageService {
	return ImageService{
		client: client,
	}
}

func (s *ImageService) CreateProductImage(productId int, filePath string) error {
	queryParams := url.Values{}
	queryParams.Add("image", filePath)

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	path := fmt.Sprintf(imagesProductBasePath, productId)
	if err := s.client.FormDataPost(path, queryParams, fileReader); err != nil {
		return err
	}

	return nil
}
