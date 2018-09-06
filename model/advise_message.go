package model

import advise "changelog/domain"

type AdviseMessageData struct {
	Code        string
	Type        string
	Description string
	Level       AdviseMessageLevel
	External    *advise.ExternalMessage
}

// AdviseMessageLevel
type AdviseMessageLevel string

const (
	AdviseMessageLevelWarn  = AdviseMessageLevel("WARN")
	AdviseMessageLevelERROR = AdviseMessageLevel("ERROR")
	AdviseMessageLevelINFO  = AdviseMessageLevel("INFO")
)
