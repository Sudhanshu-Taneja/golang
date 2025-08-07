package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("missing API key")
	}
	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid API key format")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid API key prefix, expected 'ApiKey'")
	}

	return vals[1], nil
}
