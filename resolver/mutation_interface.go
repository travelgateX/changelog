package resolver

import (
	"changelog/model"
)

// Mutation interaface functions
type Mutation interface {
	CreateChange(input struct{ Input *model.CreateChangeInput }) *ChangeResolver
	UpdateChange(input struct{ Input *model.UpdateChangeInput }) *ChangeResolver
	DeleteChange(input struct{ Input *model.DeleteChangeInput }) *ChangeResolver
}
