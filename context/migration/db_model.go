package main

import (
	"changelog/scalars/datetime"
)

// Author :
type Author struct {
	ID        uint              `gorm:"type:serial"`
	Name      string            `gorm:"type: varchar(255)"`
	CreatedAt datetime.DateTime `gorm:"type:timestamp"`
	UpdatedAt datetime.DateTime `gorm:"type:timestamp"`
}

// Source :
type Source struct {
	ID        uint   `gorm:"type:serial"`
	Platform  string `gorm:"type:varchar(20)"`
	AuthorID  uint
	CreatedAt datetime.DateTime `gorm:"type:timestamp"`
	UpdatedAt datetime.DateTime `gorm:"type:timestamp"`
}
