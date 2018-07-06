package model

import (
	"changelog/scalars/datetime"
	"changelog/scalars/version"

	graphql "github.com/graph-gophers/graphql-go"
)

// Release :
type Release struct {
	ReleaseData *ReleaseData
}

// ReleaseData :
type ReleaseData struct {
	Code        graphql.ID
	Version     version.Version
	Name        *string
	ReleaseDate datetime.DateTime
	CreatedAt   datetime.DateTime
	UpdatedAt   datetime.DateTime
}
