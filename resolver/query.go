package resolver

import (
	"context"
	"graphql-go-example/errors"

	"github.com/jinzhu/gorm"
)

// QueryResolver : is the entry point for all top-level read operations.
type QueryResolver struct {
	db *gorm.DB
}

// NewRoot : generates root entry point
func NewRoot(db *gorm.DB) (*QueryResolver, error) {
	if db == nil {
		// ========= TODO: handle custom errors
		return nil, errors.New("UnableToResolve")
	}

	return &QueryResolver{db: db}, nil
}

// ChangeTypesQueryArgs : are the arguments for the "ChangeTypes" query.
type ChangeTypesQueryArgs struct {
}

// ChangeTypes : resolves a list of all change types
func (r QueryResolver) ChangeTypes(ctx context.Context, args ChangeTypesQueryArgs) (*[]*ChangeTypeResolver, error) {
	// ========= TODO: return something nice
	return nil, nil
}
