package main

import (
	"changelog/config"
	"changelog/context"
	"changelog/handler"
	"changelog/resolver"
	"changelog/schema"

	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

func main() {
	log.Printf("====== Changelog server ON ======")

	config := config.MustLoadConfig()                  // load config file
	db := context.MustOpenDB(config.GetDBConnString()) // load data base
	defer db.Close()                                   // defer database close

	qRoot := resolver.NewQueryRoot(db) // get query root resolver
	//mRoot := resolver.NewMutationRoot(db)     // get mutation root resolver
	schema := schema.String(config.DebugMode) // full schema string

	// Attach parsed schema to handler.
	h := handler.GraphQL{
		Schema: graphql.MustParseSchema(schema, qRoot),
	}

	// Register handlers to routes.
	r := http.NewServeMux()
	r.Handle("/", handler.GraphiQL{})
	r.Handle("/graphql", h)
	r.Handle("/graphql/", h)

	// Configure the HTTP server.
	s := &http.Server{
		Handler:           r,
		Addr:              ":" + config.HTTPPort,
		ReadHeaderTimeout: config.HTTPReadHeaderTimeout,
		WriteTimeout:      config.HTTPWriteTimeout,
		IdleTimeout:       config.HTTPIdleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	// Start HTTP server.
	log.Printf("Listening for requests at %s", config.HTTPAddr())
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Changelog server has exploited: %s", err)
	}

	log.Printf("====== Changelog server OFF ======")
}
