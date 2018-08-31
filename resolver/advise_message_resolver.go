package resolver

import (
	"changelog/domain"
	"changelog/model"

	graphql "github.com/graph-gophers/graphql-go"
)

// AdviseMessageResolver for AdviseMessgae
type AdviseMessageResolver struct {
	AdviseMessage *model.AdviseMessageData
}

// Code :
func (r *AdviseMessageResolver) Code() graphql.ID {
	return graphql.ID(r.AdviseMessage.Code)
}

// Type :
func (r *AdviseMessageResolver) Type() string {
	return r.AdviseMessage.Type
}

//Description :
func (r *AdviseMessageResolver) Description() string {
	return r.AdviseMessage.Description
}

// Level :
func (r *AdviseMessageResolver) Level() model.AdviseMessageLevel {
	return r.AdviseMessage.Level
}

// External :
func (r *AdviseMessageResolver) External() *ExternalMessageResolver {
	return &ExternalMessageResolver{r.AdviseMessage.External}
}

// ExternalMessageResolver resolver for ExternalMessage
type ExternalMessageResolver struct {
	ExternalMessage *domain.ExternalMessage
}

// Code : externam message code
func (r *ExternalMessageResolver) Code() *string {
	return r.ExternalMessage.Code
}

// Message : external message text
func (r *ExternalMessageResolver) Message() string {
	return r.ExternalMessage.Message
}
