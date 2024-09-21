package server

import (
	"htmx/cmd/web"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestHandler(t *testing.T) {
	server := httptest.NewServer(templ.Handler(web.Home()))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}

	expected := "Get Random Number"
	if !strings.Contains(string(body), expected) {
		t.Errorf("expected response body to contain %v; got %v", expected, string(body))
	}
}
