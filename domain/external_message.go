package domain

// ExternalMessage:
type ExternalMessage struct {
	Code    *string `json:"code"`
	Message string  `json:message`
}
