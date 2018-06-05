package model

// Commit : commit data model
type Commit struct {
	ID        int32    `json:"id" gorm:"type:serial"`
	Message   string   `json:"message" gorm:"type:text"`
	User      string   `json:"user"`
	Release   string   `json:"release"`
	Resource  string   `json:"resource"`
	Category  Category `json:"category" gorm:"type:varchar(10)"`
	CreatedAt string   `json:"created_at"`
}
