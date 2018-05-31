package model

// Commit : data model
type Commit struct {
	ID        string `json:"id" db:"id"`
	Message   string `json:"message" db:"message"`
	User      string `json:"user" db:"user"`
	Release   string `json:"release" db:"release"`
	Resource  string `json:"resource" db:"resource"`
	Category  string `json:"category" db:"category"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
