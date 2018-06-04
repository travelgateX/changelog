package model

// Commit : commit data model
type Commit struct {
	ID        string   `json:"id" db:"id" sql:"type:integer;primary_key"`
	Message   string   `json:"message" db:"message" sql:"type:text"`
	User      string   `json:"user" db:"user"`
	Release   string   `json:"release" db:"release"`
	Resource  string   `json:"resource" db:"resource"`
	Category  Category `json:"category" db:"category" sql:"type:varchar(10)"`
	CreatedAt string   `json:"created_at" db:"created_at"`
}
