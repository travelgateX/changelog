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
	"github.com/jinzhu/gorm"
)

func main() {
	log.Printf("====== Changelog server ON ======")
	defer log.Printf("====== Changelog server OFF ======")

	var (
		err error
		c   *config.Config
		db  *gorm.DB
		sch *graphql.Schema
	)

	// Load config file
	if c, err = config.LoadConfig("./config"); err != nil {
		log.Fatalf("fatal error, cant parse config file. %v", err)
	}

	// Open database connection
	if db, err = context.OpenDB(c); err != nil {
		log.Fatalf("fatal error, cant open database connection. %v", err)
	}
	defer db.Close()

	// Parse graphql schema
	if sch, err = graphql.ParseSchema(schema.String(c.DebugMode), resolver.NewRoot(db)); err != nil {
		log.Fatalf("fatal error, cant parse graphql schema. %+v", err)
	}

	// Attach parsed schema to handler.
	h := handler.GraphQL{Schema: sch}

	// Register handlers to routes.
	r := http.NewServeMux()
	r.Handle("/", handler.GraphiQL{})
	r.Handle("/graphql", h)
	r.Handle("/graphql/", h)

	// Configure the HTTP server.
	s := &http.Server{
		Handler:           r,
		Addr:              ":" + c.HTTPPort,
		ReadHeaderTimeout: c.HTTPReadHeaderTimeout,
		WriteTimeout:      c.HTTPWriteTimeout,
		IdleTimeout:       c.HTTPIdleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	// Start HTTP server.
	log.Printf("Listening for requests at %s", c.HTTPAddr())
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("fatal error, changelog server has exploited: %s", err)
	}

}
