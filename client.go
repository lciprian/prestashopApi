package prestashopApi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/lciprian/prestashopApi/models"
)

type Prestashop struct {
	XMLName      xml.Name                  `xml:"prestashop"`
	Product      *models.ProductRequest    `xml:"product,omitempty"`
	Combinations *models.ProductVariantReq `xml:"combination,omitempty"`
}

type Prestashop2 struct {
	XMLName     xml.Name           `xml:"prestashop"`
	Product     models.Product     `xml:"product,omitempty"`
	Combination models.Combination `xml:"combination,omitempty"`
}

type Client struct {
	client  *http.Client
	version string
	baseURL *url.URL
	apiKey  string
}

func (c *Client) NewRequest(method, relPath string, params url.Values, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("api/%s", relPath))
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	cType := "application/json"
	if method == "POST" {
		cType = "application/xml"
	}
	req.Header.Add("Content-Type", cType)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Output-Format", "JSON")
	req.Header.Add("Io-Format", "JSON")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.apiKey))

	return req, nil
}

func (c *Client) NewFormDataRequest(method, relPath string, params url.Values, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("api/%s", relPath))
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Output-Format", "JSON")
	req.Header.Add("Io-Format", "JSON")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.apiKey))

	return req, nil
}

func (c *Client) Get(path string, params url.Values, resource interface{}) error {
	req, err := c.NewRequest("GET", path, params, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkResponseEmptyOrError(resp); err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//	fmt.Printf("Body : %s", body)

	if err := json.Unmarshal(body, resource); err != nil {
		//check for empty list
		var list []interface{}
		if err2 := json.Unmarshal(body, &list); err2 == nil {
			return nil
		}

		return err
	}

	return nil
}

func (c *Client) FormDataPost(path string, params url.Values, resource io.Reader) error {
	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form field for the image file
	part, err := writer.CreateFormFile("image", params.Get("image"))
	if err != nil {
		return fmt.Errorf("error creating form file: %s", err)
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, resource)
	if err != nil {
		return fmt.Errorf("error copying file to form: %s", err.Error())
	}
	fmt.Println("FormDataContentType----------", writer.FormDataContentType())

	// Close the multipart writer to finalize the request body
	writer.Close()

	req, err := c.NewFormDataRequest("POST", path, params, &requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := c.checkResponseEmptyOrError(resp); err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Body : %d", len(body))

	return nil
}

func (c *Client) Post(path string, params url.Values, resource io.Reader) ([]byte, error) {
	req, err := c.NewRequest("POST", path, params, resource)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkResponseEmptyOrError(resp); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Body : %s", body)

	return body, nil
}

func (c *Client) checkResponseEmptyOrError(r *http.Response) error {
	if http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("response message error: %s", err.Error())
	}

	respError := fmt.Errorf("response error status: %d", r.StatusCode)
	errors := models.ResponseErrors{}
	if err := json.Unmarshal(data, &errors); err != nil {
		return respError
	}

	if len(errors.Errors) > 0 {
		respError = fmt.Errorf("response error code: %d, message: %s", errors.Errors[0].Code, errors.Errors[0].Message)
	}

	return respError
}

//type T struct {
//	Combinations []struct {
//		Id                string      `json:"id"`
//		IdProduct         string      `json:"id_product"`
//		Location          interface{} `json:"location"`
//		Ean13             interface{} `json:"ean13"`
//		Isbn              interface{} `json:"isbn"`
//		Upc               interface{} `json:"upc"`
//		Mpn               interface{} `json:"mpn"`
//		Quantity          interface{} `json:"quantity"`
//		Reference         interface{} `json:"reference"`
//		SupplierReference interface{} `json:"supplier_reference"`
//		WholesalePrice    interface{} `json:"wholesale_price"`
//		Price             string      `json:"price"`
//		Ecotax            interface{} `json:"ecotax"`
//		Weight            string      `json:"weight"`
//		UnitPriceImpact   interface{} `json:"unit_price_impact"`
//		MinimalQuantity   string      `json:"minimal_quantity"`
//		LowStockThreshold interface{} `json:"low_stock_threshold"`
//		LowStockAlert     interface{} `json:"low_stock_alert"`
//		DefaultOn         interface{} `json:"default_on"`
//		AvailableDate     interface{} `json:"available_date"`
//		Associations      struct {
//			ProductOptionValues []struct {
//				Id string `json:"id"`
//			} `json:"product_option_values"`
//			Images []struct {
//				Id string `json:"id"`
//			} `json:"images"`
//		} `json:"associations"`
//	} `json:"combinations"`
//}
