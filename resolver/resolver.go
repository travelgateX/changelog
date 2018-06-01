package resolver

import (
	"github.com/jinzhu/gorm"
)

// Resolver : root resolver
type Resolver struct{}

// NewQueryRoot : generates query root entry point
func NewQueryRoot(db *gorm.DB) *QueryResolver {
	return &QueryResolver{db: db}
}

// NewMutationRoot : generates mutation root entry point
func NewMutationRoot(db *gorm.DB) *MutationResolver {
	return &MutationResolver{db: db}
}
