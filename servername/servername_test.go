package servername

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServerNameMiddleware is a unit test function that tests the behavior of the ServerNameMiddleware.
// It verifies that the middleware sets the correct server name in the response header.
// The function creates a test HTTP handler, sets the server name, and makes a test request.
// It then checks if the response status code and server header match the expected values.
func TestServerNameMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serverName := "TestServer"
	middleware := ServerName(serverName)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	serverHeader := rr.Header().Get("Server")
	if serverHeader != serverName {
		t.Errorf("Server header mismatch: got %v want %v",
			serverHeader, serverName)
	}
}
