package prestashopApi

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func (s *ImageService) CreateProductImage(productId int, imageUrl string) error {
	queryParams := url.Values{}
	queryParams.Add("image", imageUrl)

	// Open the image file
	fileReader, err := s.getAsset(imageUrl)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(imagesProductBasePath, productId)
	if err := s.client.FormDataPost(path, queryParams, fileReader); err != nil {
		return err
	}

	return nil
}

func (s *ImageService) getAsset(assetUrl string) (io.Reader, error) {
	response, err := http.Get(assetUrl)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response code: %d", response.StatusCode))
	}

	respByte, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(respByte), nil
}
