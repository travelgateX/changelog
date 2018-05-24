package tools

import (
	"net/http"
)

// Path : get absolute path of a request
func Path(r *http.Request) string {
	return r.Host + r.URL.Path
}
