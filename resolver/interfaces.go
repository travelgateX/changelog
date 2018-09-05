package resolver

import (
	"changelog/config"
	"changelog/model"

	"github.com/jinzhu/gorm"
)

// Mutations interface declaration

// InputCreates args
type InputCreates struct {
	Input *model.CreateChangeInput
}

// InputUpdates args
type InputUpdates struct {
	Input *model.UpdateChangeInput
}

// InputDeletes args
type InputDeletes struct {
	Input *model.DeleteChangeInput
}

// Mutations interface
type Mutations interface {
	CreateChange(InputCreates) *ChangeResolver
	DeleteChange(InputUpdates) *ChangeResolver
	UpdateChange(InputDeletes) *ChangeResolver
}

// DBinterface interface declaration
type DBinterface interface {
	OpenDB(c *config.Config) (*gorm.DB, error)
	MustOpen(c *config.Config) *gorm.DB
}
