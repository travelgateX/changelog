package resolver

import (
	"changelog/config"
	"changelog/model"

	"github.com/jinzhu/gorm"
)

// Mutation interaface functions
type Mutation interface {
	CreateChange(input struct{ Input *model.CreateChangeInput }) *ChangeResolver
	UpdateChange(input struct{ Input *model.UpdateChangeInput }) *ChangeResolver
	DeleteChange(input struct{ Input *model.DeleteChangeInput }) *ChangeResolver
}

type DataBase interface {
	OpenDB(c *config.Config) (*gorm.DB, error)
	MustOpenDB(c *config.Config) *gorm.DB
}
