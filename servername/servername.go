package servername

import "net/http"

// ServerName is a middleware that adds the specified server name to the response header.
// It takes a name string as input and returns a function that wraps the provided http.Handler.
// The returned function is used as middleware to set the "Server" header in the HTTP response.
func ServerName(name string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", name)
			next.ServeHTTP(w, r)
		})
	}
}
