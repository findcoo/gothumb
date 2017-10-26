package gothumb

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"
)

var (
	baseURL = &url.URL{
		Scheme: "https",
		Host:   "api.bithumb.com",
	}
)

// APIResponse API 공통 응답
type APIResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

// Client API client
type Client struct {
	info                *url.URL
	personalToken       string
	personalTokenSecret string
	*http.Client
}

// NewClient client 생성자
func NewClient(pToken, pTokenSecret string) *Client {
	info := *baseURL
	info.Path = "info"

	return &Client{
		info:                &info,
		personalToken:       pToken,
		personalTokenSecret: pTokenSecret,
		Client:              &http.Client{Timeout: time.Second * 10},
	}
}

func (c *Client) doInfo(uriPath string, formValue url.Values) (*APIResponse, error) {
	reqURL := *c.info
	reqURL.Path = path.Join(reqURL.Path, uriPath)

	formValue.Add("apiKey", c.personalToken)
	formValue.Add("secretKey", c.personalTokenSecret)

	resp, err := c.PostForm(reqURL.String(), formValue)
	if err != nil {
		return nil, err
	}
	unmarshal := json.NewDecoder(resp.Body)

	apiResp := &APIResponse{}
	if err := unmarshal.Decode(apiResp); err != nil {
		return nil, err
	}
	return apiResp, nil
}

// Balance 사용자의 잔고 조회
func (c *Client) Balance(currency string) (*APIResponse, error) {
	return c.doInfo("balance", url.Values{"currency": []string{currency}})
}

// DepositAddress 사용자의 지갑 조회
func (c *Client) DepositAddress(currency string) (*APIResponse, error) {
	return c.doInfo("wallet_address", url.Values{"currency": []string{currency}})
}
