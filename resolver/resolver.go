package resolver

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Resolver : root resolver
type Resolver struct{}

// NewRoot : generates root entry point
// ==========  <necessary explanation>  ==========
func NewRoot(db *gorm.DB) (*QueryResolver, error) {
	if db == nil {
		// ========= TODO: handle custom errors
		return nil, errors.New("UnableToResolve")
	}
	return &QueryResolver{db: db}, nil
} // ==========  <necessary explanation>  ========
