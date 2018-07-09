package main

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Platform :
type Platform string

// Source :
type Source struct {
	gorm.Model
	Code     graphql.ID `gorm:"type:varchar(100)"`
	Platform Platform   `gorm:"type:varchar(20)"`
	Author   string     `gorm:"type:varchar(255)"`
	Changes  []Change
}

// Version :
type Version string

// Release :
type Release struct {
	gorm.Model
	Code        graphql.ID `gorm:"type:varchar(100)"`
	Name        string     `gorm:"type:varchar(255)"`
	Version     Version    `gorm:"type:varchar(255)"`
	ReleaseDate time.Time  `gorm:"type:timestamp"`
	Changes     []Change
}

// Change :
type Change struct {
	gorm.Model
	Code       graphql.ID       `gorm:"type:varchar(100)"`
	Message    string           `gorm:"type:text"`
	Project    *string          `gorm:"type:varchar(50)"`
	ChangeDate time.Time        `gorm:"type:timestamp"`
	Type       model.ChangeType `gorm:"type:varchar(20)"`
	ReleaseID  uint
	SourceID   uint
}
