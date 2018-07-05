package model

import (
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// Release :
type Release struct {
	Code        graphql.ID        `db:"code_id" gorm:"type:serial"`
	Version     string            `db:"version"`
	Name        *string           `db:"name"`
	ReleaseDate datetime.DateTime `db:"release_date"`
	CreatedAt   datetime.DateTime `db:"created_at"`
	UpdatedAt   datetime.DateTime `db:"updated_at"`
}
