package auth

import (
	"net/url"
)

type Service struct {
	APIKey  string
	BaseURL string
}

func (s *Service) BuildAuthURL(token string) (string, error) {
	u, err := url.Parse("https://www.last.fm/api/auth")
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("api_key", s.APIKey)
	q.Set("token", token)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
