package prestashopApi

var productBasePath = "products"

type ProductService struct {
	client *Client
}

type Product struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductList struct {
	Products []Product `json:"products,omitempty"`
}

func newProductService(client *Client) ProductService {
	return ProductService{
		client: client,
	}
}

func (s *ProductService) ListProducts() ([]Product, error) {
	productList := ProductList{}

	err := s.client.Get(productBasePath, nil, &productList)
	if err != nil {
		return nil, err
	}

	return productList.Products, err
}
