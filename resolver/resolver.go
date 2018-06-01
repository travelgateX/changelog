package resolver

import (
	"github.com/jinzhu/gorm"
)

// Resolver : root resolver
type Resolver struct {
	db *gorm.DB
}

// NewRoot : generates query root entry point
func NewRoot(db *gorm.DB) *Resolver {
	return &Resolver{db: db}
}
