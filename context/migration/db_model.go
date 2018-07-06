package main

import (
	"changelog/scalars/datetime"

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
}

// Version :
type Version string

// Release :
type Release struct {
	gorm.Model
	Code        graphql.ID        `gorm:"type:varchar(100)"`
	Name        string            `gorm:"type:varchar(255)"`
	Version     Version           `gorm:"type:varchar(255)"`
	releaseDate datetime.DateTime `gorm:"type:timestamp"`
}
