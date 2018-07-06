package main

import (
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
