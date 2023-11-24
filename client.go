package prestashopApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Client struct {
	client  *http.Client
	version string
	baseURL *url.URL
	apiKey  string
}

func (c *Client) NewRequest(method, relPath string, params url.Values, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("api/%s", relPath))
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	var js []byte = nil
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
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

	fmt.Printf("Body : %s", body)

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

func (c *Client) Post(path string, params url.Values, resource interface{}) error {
	req, err := c.NewRequest("POST", path, params, resource)
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

	fmt.Printf("Body : %s", body)

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

func (c *Client) checkResponseEmptyOrError(r *http.Response) error {
	if http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	if _, err := io.ReadAll(r.Body); err != nil {
		return fmt.Errorf("response message error: %s", err.Error())
	}

	return fmt.Errorf("response error status: %d", r.StatusCode)
}
