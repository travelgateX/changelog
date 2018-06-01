package main

import (
	"changelog/config"
	"changelog/handler"
	"changelog/resolver"
	"changelog/schema"
	"changelog/service"

	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

func main() {
	log.Printf("====== Changelog server ON ======")

	// Load environment variables.
	config := config.MustLoadConfig()

	// Get database connection
	db := service.MustOpenDB(config.GetDBConnString())

	// Get root resolver
	root, err := resolver.NewRoot(db)
	if err != nil {
		log.Fatalf("cant get root resolver: %s", err)
	}

	// Attach parsed schema to handler.
	h := handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.String(config.DebugMode), root),
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
