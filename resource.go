package prestashopApi

import (
	"fmt"
)

var resourceBasePath = ""

type ResourceService struct {
	client *Client
}

type Resource struct {
	Name string `json:"name,omitempty"`
}

func newResourceService(client *Client) ResourceService {
	return ResourceService{
		client: client,
	}
}

func (s *ResourceService) ListResources() ([]Resource, error) {
	resource := make([]Resource, 0)
	err := s.client.Get(resourceBasePath, &resource)
	if err != nil {
		return nil, err
	}

	fmt.Println("-----------", resource)

	return resource, err
}
