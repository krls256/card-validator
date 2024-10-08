package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *Client) RequestWithBody(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequestWithContext(ctx, method, c.GetURL(url), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) NewPostRequest(ctx context.Context, url string, body interface{}) (*http.Request, error) {
	return c.RequestWithBody(ctx, http.MethodPost, url, body)
}

func (c *Client) GetURL(path string) string {
	baseURL := strings.TrimRight(c.cfg.Path(), "/")

	return baseURL + path
}

func HandleResponse[T any](req *http.Request) (T, error) {
	var defaultResult T

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return defaultResult, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultResult, err
	}

	if err = json.Unmarshal(respBody, &defaultResult); err != nil {
		return defaultResult, err
	}

	return defaultResult, nil
}
