package tclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req/v3"
	"strconv"
	"time"
)

type Client struct {
	*req.Client
	service string
}

func (c *Client) WithRegion(region string) *Client {
	c.SetCommonHeaderNonCanonical("X-TC-Region", region)
	return c
}

func (c *Client) WithCredential(secretId, secretKey string) *Client {
	c.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		requestTimestamp := req.Headers["X-TC-Timestamp"][0]
		ct := req.Headers.Get("Content-Type")
		canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", ct, fmt.Sprintf("%s.tencentcloudapi.com", c.service))
		signedHeaders := "content-type;host"
		hashedRequestPayload := sha256hex(req.Body)
		canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
			"POST",
			"/",
			"",
			canonicalHeaders,
			signedHeaders,
			hashedRequestPayload,
		)
		// build string to sign
		algorithm := "TC3-HMAC-SHA256"
		timestamp, _ := strconv.ParseInt(requestTimestamp, 10, 64)
		t := time.Unix(timestamp, 0).UTC()
		// must be the format 2006-01-02, ref to package time for more info
		date := t.Format("2006-01-02")
		credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, c.service)
		hashedCanonicalRequest := sha256hex([]byte(canonicalRequest))
		string2sign := fmt.Sprintf("%s\n%s\n%s\n%s",
			algorithm,
			requestTimestamp,
			credentialScope,
			hashedCanonicalRequest)

		// sign string
		secretDate := hmacsha256(date, "TC3"+secretKey)
		secretService := hmacsha256(c.service, secretDate)
		secretKey := hmacsha256("tc3_request", secretService)
		signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretKey)))

		// build authorization
		authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
			algorithm,
			secretId,
			credentialScope,
			signedHeaders,
			signature,
		)
		req.SetHeader("Authorization", authorization)
		return nil
	})
	return c
}

func sha256hex(b []byte) string {
	bb := sha256.Sum256(b)
	return hex.EncodeToString(bb[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func NewClient(service, version string) *Client {
	client := &Client{
		Client:  req.NewClient(),
		service: service,
	}
	client.SetBaseURL(fmt.Sprintf("https://%s.tencentcloudapi.com", service))
	if version != "" {
		client.SetCommonHeaderNonCanonical("X-TC-Version", version)
	}
	return client
}
