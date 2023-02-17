package evotorrestogo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ProdURL = "https://resto.evotor.tech/api/3rdparty"
	TestURL = "https://resto-test.evotor.tech/api/3rdparty"
	DevURL  = "https://resto-dev.evotor.tech/api/3rdparty"
)

// Клиент подключения к API
type Client struct {
	BaseURL string
	Token   string
	Http    *http.Client
}

// Получить список торговых точек
func (c *Client) SelectStores(ctx context.Context) ([]Store, error) {
	url := c.getUrl("/store")

	stores := []Store{}

	return stores, c.doRequest(ctx, "GET", url, nil, nil, &stores)
}

// Получить меню
func (c *Client) SelectMenu(ctx context.Context, storeId string) ([]MenuItem, error) {
	url := c.getUrl("/product/" + storeId)

	products := []MenuItem{}

	return products, c.doRequest(ctx, "GET", url, nil, nil, &products)
}

// Создать заказ
func (c *Client) CreateOrder(ctx context.Context, storeId string, order Order) (Order, error) {
	url := c.getUrl("/order/" + storeId)

	result := Order{}

	return result, c.doRequest(ctx, "POST", url, nil, order, &result)
}

// Получить заказ по ИД
func (c *Client) GetOrder(ctx context.Context, storeId string, orderId string) (Order, error) {
	url := c.getUrl("/order/" + storeId + "/" + orderId)

	result := Order{}

	return result, c.doRequest(ctx, "GET", url, nil, nil, &result)
}

func (c *Client) getUrl(path string) string {
	if c.BaseURL != "" {
		return c.BaseURL + path
	}
	return ProdURL + path
}

func (c *Client) doRequest(ctx context.Context, method, url string, headers map[string]string, body any, response any) error {
	var payload io.Reader = nil
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("can't serialize body: %w", err)
		}
		payload = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return fmt.Errorf("can't create request: %w", err)
	}
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Accept", "application/json")
	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpClient := c.Http
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected http code: %s %s", resp.Status, string(body))
	}

	if response == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("can't parse response: %w", err)
	}

	return nil
}
