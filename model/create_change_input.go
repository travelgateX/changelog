package model

// CreateChangeInput :
type CreateChangeInput struct {
	Code    *string
	Message string
	Project string
	Type    ChangeType
}
