# ServerName Middleware

The `servername` package provides a middleware function called `ServerName` that adds a specified server name to the response header of HTTP requests.

Usage
-----

### Import 

```zsh
go get github.com/ugurkorkmaz/middleware/servername
```

### Parameters

-   `name`: A string representing the server name to be added to the response header.

### Return Value

The `ServerName` function returns a function that wraps the provided `http.Handler`.

### Example

```go
package main

import (
	"net/http"
	"github.com/ugurkorkmaz/servername"
)

func main() {
	// Create a new HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Set the server name to be added to the response header
	serverName := "MyCustomServer"

	// Wrap the HTTP handler with the ServerName middleware
	http.Handle("/", servername.ServerName(serverName)(handler))

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
```

In this example, the `ServerName` middleware is used to set the "Server" header of the HTTP response to "MyCustomServer".
