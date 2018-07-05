package model

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Source :
type Source struct {
	Code     graphql.ID `db:"code_id" gorm:"type:serial"`
	Platform SourcePlatform
	Author   Author
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
