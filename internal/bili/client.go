package bili

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type BiliClient struct {
	httpClient *http.Client
	signer     *WbiSigner
	cookie     string
}

type BiliResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func NewBiliClient() *BiliClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &BiliClient{
		httpClient: httpClient,
		signer:     NewWbiSigner(httpClient),
		cookie:     os.Getenv("BILI_COOKIE"),
	}
}

func (c *BiliClient) doRequest(method, path string, params map[string]string, needWbi bool) ([]byte, error) {
	var url string
	if needWbi {
		url = "https://api.bilibili.com" + path + "?" + c.signer.Sign(params)
	} else {
		// Simple query string for non-WBI
		query := ""
		for k, v := range params {
			if query != "" {
				query += "&"
			}
			query += fmt.Sprintf("%s=%s", k, v)
		}
		url = "https://api.bilibili.com" + path + "?" + query
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Inject Cookie
	if c.cookie != "" {
		req.Header.Set("Cookie", c.cookie)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var biliResp BiliResponse
	if err := json.Unmarshal(body, &biliResp); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w, body: %s", err, string(body))
	}

	if biliResp.Code != 0 {
		return nil, fmt.Errorf("bilibili api error: code=%d, message=%s", biliResp.Code, biliResp.Message)
	}

	return biliResp.Data, nil
}

func (c *BiliClient) doRawRequest(method, path string, params map[string]string, needWbi bool) ([]byte, error) {
	var url string
	if needWbi {
		url = "https://api.bilibili.com" + path + "?" + c.signer.Sign(params)
	} else {
		query := ""
		for k, v := range params {
			if query != "" {
				query += "&"
			}
			query += fmt.Sprintf("%s=%s", k, v)
		}
		url = "https://api.bilibili.com" + path + "?" + query
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.cookie != "" {
		req.Header.Set("Cookie", c.cookie)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
