package prestashopApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	if err := c.checkResponseError(resp); err != nil {
		return err
	}

	if body, err := io.ReadAll(resp.Body); err == nil {
		fmt.Printf("Body : %s", body)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&resource); err != nil {
		return err
	}

	return nil
}

func (c *Client) checkResponseError(r *http.Response) error {
	if http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	if _, err := io.ReadAll(r.Body); err != nil {
		return fmt.Errorf("response message error: %s", err.Error())
	}

	return fmt.Errorf("response error status: %d", r.StatusCode)
}
