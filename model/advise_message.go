package model

// AdviseMessageData :
type AdviseMessageData struct {
	Code        string
	Type        string
	Description string
	Level       AdviseMessageLevel
	External    *ExternalMessage
}

// AdviseMessageLevel :
type AdviseMessageLevel string

// AdviseMessageLevelWarn :
const (
	AdviseMessageLevelWarn  = AdviseMessageLevel("WARN")
	AdviseMessageLevelERROR = AdviseMessageLevel("ERROR")
	AdviseMessageLevelINFO  = AdviseMessageLevel("INFO")
)

// ExternalMessage :
type ExternalMessage struct {
	Code    *string `json:"code"`
	Message string  `json:"message"`
}
