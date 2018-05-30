package handler

import (
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

// GraphQL : handles GraphQL API requests over HTTP.
type GraphQL struct {
	Schema *graphql.Schema
}

// query : represents a single GraphQL query.
type query struct {
	OpName    string                 `json:"operationName"`
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// ========= <different solutions found> =========
// ServeHTTP :
func (h GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validate the request.
	if ok := isSupported(r.Method); !ok {
		respond(w, errorJSON("only POST or GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}

	var q query
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		respond(w, errorJSON(err.Error()), http.StatusBadRequest)
		return
	}

	// ========= TODO: authentification here?
	// ==========  <necessary explanation>  ==========
	// ==> authorization, company, context??????
	// company := r.Context().Value(authorization.ContextKey)
	// ctx := c.WithValue(r.Context(), authorization.ContextKey, nil)
	res := h.Schema.Exec(r.Context(), q.Query, q.OpName, q.Variables)

	// Encode json graphql response
	resJSON, err := json.Marshal(res)
	if err != nil {
		respond(w, errorJSON("server error"), http.StatusInternalServerError)
		return
	}

	respond(w, resJSON, http.StatusOK)
}
