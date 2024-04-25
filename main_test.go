package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEndpointHandler(t *testing.T) {

	tests := []struct {
		name       string
		query      string
		statuscode int
		body       string
	}{
		{
			name:       "valid message",
			query:      "message=goodmorning",
			statuscode: 200,
			body:       "goodmorning",
		},
		{
			name:       "no message",
			query:      "",
			statuscode: 400,
			body:       "No message provided",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/echo?"+test.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(EndpointHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.statuscode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.statuscode)
			}

			if body := strings.TrimSpace(rr.Body.String()); body != test.body {
				t.Errorf("handler returned unexpected body: got %v want %v", body, test.body)
			}
		})
	}
}
