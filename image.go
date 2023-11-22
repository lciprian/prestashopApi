package prestashopApi

import (
	"bytes"
	"fmt"
	"net/url"
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

func (s *ImageService) CreateProductImage(productId int, requestBody bytes.Buffer) error {
	queryParams := url.Values{}

	path := fmt.Sprintf(imagesProductBasePath, productId)
	if err := s.client.FormDataPost(path, queryParams, &requestBody); err != nil {
		return err
	}

	return nil
}
