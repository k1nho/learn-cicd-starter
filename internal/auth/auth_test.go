package auth

import (
	"errors"
	"net/http"
	"testing"
)

type result struct {
	apikey string
	err    error
}

func compareResults(r1 result, r2 result) bool {
	return (r1.apikey == r2.apikey && errors.Is(r1.err, r2.err))
}

func TestGetApiKey(t *testing.T) {

	header := http.Header{
		"Authorization": []string{"ApiKey key"},
	}
	invalidHeader := http.Header{}
	noAuthValue := http.Header{
		"Authorization": []string{"ApiKey"},
	}

	tests := map[string]struct {
		input http.Header
		want  result
	}{
		"simple":                    {input: header, want: result{apikey: "key", err: nil}},
		"no auth header":            {input: invalidHeader, want: result{apikey: "", err: ErrNoAuthHeaderIncluded}},
		"auth header with no value": {input: noAuthValue, want: result{apikey: "", err: ErrMalformedHeader}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.input)
			got := result{apikey: key, err: err}

			if !compareResults(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
