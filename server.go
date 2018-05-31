package main

import (
	"changelog/config"
	"changelog/handler"
	"changelog/model"
	"changelog/resolver"
	"changelog/schema"

	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

func main() {
	log.Printf("====== Changelog server ON ======")
	defer log.Printf("====== Changelog server OFF ======")

	// Load environment variables.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cant load config file: %s", err)
	}

	// Get database connection
	db := model.MustOpenDB(config.GetConnString())

	// Get root resolver
	root, err := resolver.NewRoot(db)
	if err != nil {
		log.Fatalf("cant get root resolver: %s", err)
	}

	sch, err := graphql.ParseSchema(schema.String(), root)
	if err != nil {
		log.Fatalf("su puta malder: %v", err)
	}

	// Create the request handler and parse schema.
	h := handler.GraphQL{
		Schema: sch,
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
}
