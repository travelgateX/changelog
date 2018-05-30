package main

import (
	"changelog/config"
	"changelog/handler"

	"log"
	"net/http"
)

func main() {
	log.Printf("====== Changelog server ON ======")
	defer log.Printf("====== Changelog server OFF ======")

	// Load environment variables.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	// ========= TODO: get root resolver
	// Create the request handler and parse schema.
	h := handler.GraphQL{
		// ========= TODO: fix load schema
		//Schema: graphql.MustParseSchema(schema.String(), nil),
	}

	// Register handlers to routes.
	router := http.NewServeMux()
	router.Handle("/graphql", h)
	router.Handle("/graphql/", h)

	// Configure the HTTP server.
	s := &http.Server{
		Handler:           router,
		Addr:              ":" + config.HTTPPort,
		ReadHeaderTimeout: config.HTTPReadHeaderTimeout,
		WriteTimeout:      config.HTTPWriteTimeout,
		IdleTimeout:       config.HTTPIdleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	// Start HTTP server.
	log.Printf("Listening for requests on port %s ==", config.HTTPAddr())
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Changelog server has exploited: %s", err)
	}
}
