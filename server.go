package main

import (
	"changelog/config"
	"changelog/handler"
	"changelog/model"
	"changelog/resolver"
	"changelog/schema"

	graphql "github.com/graph-gophers/graphql-go"

	"log"
	"net/http"
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

	// Create the request handler and parse schema.
	//str := schema.String()

	/*
		f, err := os.Create("schema.raw")
		if err != nil {
			log.Fatalf("cant write schema log file", err.Error())
		}
		if _, err := f.WriteString(str); err != nil {
			log.Fatalf("cant flush string schema in log file", err.Error())
		}
		if err := f.Sync(); err != nil {
			log.Fatalf("cant sync schema log file", err.Error())
		}
	*/

	sch, err := graphql.ParseSchema(schema.String(), root)
	if err != nil {
		log.Printf("su puta madre: %v", err)
		return
	}

	h := handler.GraphQL{
		Schema: sch,
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
	log.Printf("Listening for requests at %s", config.HTTPAddr())
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Changelog server has exploited: %s", err)
	}
}
