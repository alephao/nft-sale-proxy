package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockHttpClient(response GenericResponse) (*http.Client, string) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response.StatusCode != 0 {
			w.WriteHeader(response.StatusCode)
		}

		for key, value := range response.Headers {
			w.Header().Add(key, value)
		}

		if response.Body != "" {
			w.Write([]byte(response.Body))
		}
	}))
	return server.Client(), server.URL
}

func TestHandler(t *testing.T) {
	t.Run("no token id", func(t *testing.T) {
		response, err := HandleRequest(nil, nil, func() string { return "" })
		assertError(t, err, nil)
		assertResponse(t, response, GenericResponse{
			StatusCode: 404,
			Body:       "\"message\": \"invalid token id\"",
		})
	})

	t.Run("token id < 0", func(t *testing.T) {
		response, err := HandleRequest(&Config{}, nil, func() string { return "-1" })
		assertError(t, err, nil)
		assertResponse(t, response, GenericResponse{
			StatusCode: 404,
			Body:       "\"message\": \"id out of bounds\"",
		})
	})

	t.Run("token id > limit", func(t *testing.T) {
		response, err := HandleRequest(&Config{
			NumberOfTokens: 10,
		}, nil, func() string { return "11" })
		assertError(t, err, nil)
		assertResponse(t, response, GenericResponse{
			StatusCode: 404,
			Body:       "\"message\": \"id out of bounds\"",
		})
	})

	t.Run("happy path", func(t *testing.T) {
		mockClient, url := mockHttpClient(GenericResponse{StatusCode: 200, Body: `{"a": "b", "c": 10}`})
		config := &Config{
			NumberOfTokens: 10,
			RevealUpTo:     10,
			BaseURL:        fmt.Sprintf("%s/", url),
		}

		response, err := HandleRequest(config, mockClient, func() string { return "1" })

		assertError(t, err, nil)
		assertResponse(t, response, GenericResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"a":"b","c":10}`,
		})
	})
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("expected error\n%s\ngot\n%s", expected.Error(), got.Error())
	}
}

func assertResponse(t *testing.T, got, expected GenericResponse) {
	t.Helper()

	if got.StatusCode != expected.StatusCode || got.Body != expected.Body {
		t.Errorf("expected\n%v\ngot\n%v", expected, got)
	}

	for key, value := range got.Headers {
		if expected.Headers[key] != value {
			t.Errorf("expected\n%v\ngot\n%v", expected, got)
			return
		}
	}
}
