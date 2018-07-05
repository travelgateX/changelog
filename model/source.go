package model

import (
	"changelog/scalars/datetime"

	graphql "github.com/graph-gophers/graphql-go"
)

// Source :
type Source struct {
	Code      graphql.ID        `db:"code_id" gorm:"type:serial"`
	CreatedAt datetime.DateTime `db:"created_at"`
	UpdatedAt datetime.DateTime `db:"updated_at"`
	Author    Author
	Platform  SourcePlatform
}

// SourcePlatform : Change type
type SourcePlatform string

// Source enum
const (
	Github = SourcePlatform("GITHUB")
	Gitlab = SourcePlatform("GITLAB")
	Jira   = SourcePlatform("JIRA")
	Slack  = SourcePlatform("SLACK")
	Other  = SourcePlatform("OTHER")
)
