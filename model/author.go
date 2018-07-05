package model

import (
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// Author :
type Author struct {
	Code       graphql.ID `db:"code_id" gorm:"type:serial"`
	AuthorData AuthorData
	CreatedAt  datetime.DateTime `db:"created_at"`
	UpdatedAt  datetime.DateTime `db:"updated_at"`
}

// AuthorData :
type AuthorData struct {
	Code      graphql.ID        `db:"code_id" gorm:"type:serial"`
	Name      string            `db:"name"`
	CreatedAt datetime.DateTime `db:"created_at"`
	UpdatedAt datetime.DateTime `db:"updated_at"`
}
