package tclient

import (
	"github.com/imroc/req/v3"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type Request struct {
	*req.Request
}

func (c *Client) NewRequest(action string, params interface{}) *Request {
	u, _ := url.Parse(c.BaseURL)
	req := c.Post().SetHeadersNonCanonical(map[string]string{
		"X-TC-Action":        action,
		"X-TC-Language":      "zh-CN",
		"X-TC-RequestClient": "SDK_GO_1.0.523",
		"X-TC-Timestamp":     strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":              strconv.Itoa(rand.Int()),
		"Host":               u.Host,
	})
	if params != nil {
		req.SetBodyJsonMarshal(params)
	}
	return &Request{req}
}

func (req *Request) WithRegion(region string) *Request {
	req.SetHeader("X-TC-Region", region)
	return req
}
