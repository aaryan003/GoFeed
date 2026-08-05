package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of an HTTP request.
//Example :
// Authoization : ApiKey {insert apiKey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Authorization header is missing")
	}

	vals := strings.SplitN(val, " ", 2)
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Authorization header must start with 'ApiKey'")
	}
	return vals[1], nil
}
