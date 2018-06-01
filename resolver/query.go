package resolver

import (
	"context"

	"github.com/jinzhu/gorm"
)

// QueryResolver :
type QueryResolver struct {
	db *gorm.DB
}

// CommitsQueryArgs : commits request
type CommitsQueryArgs struct {
	Resource *string
}

// Films resolves a list of films. If no arguments are provided, all films are fetched.
func (r QueryResolver) Films(ctx context.Context, args CommitsQueryArgs) (*[]*CommitResolver, error) {
	page, err := r.client.SearchFilms(ctx, strValue(args.Title))
	if err != nil {
		return nil, err
	}

	return NewFilms(ctx, NewFilmsArgs{Page: page})
}
