package auth

import (
	"errors"
	"net/http"
	"strings"
)

// ApiKey=123
func GetAPIKey(headers http.Header) (string, error) {
	authHeaderVal := headers.Get("Auth")
	if authHeaderVal == "" {
		return "", errors.New("no auth found")
	}

	authValues := strings.Split(authHeaderVal, "=")
	authValuesLength := 2
	if len(authValues) != authValuesLength {
		return "", errors.New("malformed auth header")
	}

	if authValues[0] != "ApiKey" {
		return "", errors.New("malformed api key identifier")
	}

	return authValues[1], nil
}
