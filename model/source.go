package model

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Source :
type Source struct {
	gorm.Model
	Code     graphql.ID
	Platform SourcePlatform
	Author   string
	Changes  []Change
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
