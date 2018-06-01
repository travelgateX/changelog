package resolver

import (
	"github.com/jinzhu/gorm"
)

// Resolver : root resolver
type Resolver struct{}

// NewRoot : generates root entry point
func NewRoot(db *gorm.DB) *QueryResolver {
	return &QueryResolver{db: db}
}
