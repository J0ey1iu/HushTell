package handler

import (
	"fmt"
	"net/http"
)

// Index is just a page to let you know it's working
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>The server is running.</h1>")
}
