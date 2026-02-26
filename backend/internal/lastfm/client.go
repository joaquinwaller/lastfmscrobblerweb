package lastfm

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	APIKey    string
	APISecret string
	HTTP      *http.Client
}

func New(apiKey, apiSecret string) *Client {
	return &Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetToken(ctx context.Context) (string, error) {
	u, _ := url.Parse("https://ws.audioscrobbler.com/2.0/")
	q := u.Query()
	q.Set("method", "auth.getToken")
	q.Set("api_key", c.APIKey)
	q.Set("format", "json")

	// api_sig requerido
	apiSig := BuildAPISig(map[string]string{
		"api_key": c.APIKey,
		"method":  "auth.getToken",
	}, c.APISecret)

	q.Set("api_sig", apiSig)

	u.RawQuery = q.Encode()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("lastfm getToken failed")
	}

	var out struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if out.Token == "" {
		return "", errors.New("lastfm token empty")
	}
	return out.Token, nil
}
