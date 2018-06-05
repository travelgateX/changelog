package model

import "time"

// Commit : commit data model
type Commit struct {
	ID        int32     `json:"id" gorm:"type:serial"`
	Message   string    `json:"message" gorm:"type:text"`
	User      string    `json:"user"`
	Release   string    `json:"release"`
	Resource  string    `json:"resource"`
	Category  Category  `json:"category" gorm:"type:varchar(10)"`
	Released  bool      `json:"released"`
	CreatedAt time.Time `json:"created_at"`
}
