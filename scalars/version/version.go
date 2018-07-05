package version

import (
	"fmt"
	"strconv"
)

// Version :
type Version string

// ImplementsGraphQLType :
func (Version) ImplementsGraphQLType(name string) bool {
	return name == "Version"
}

// UnmarshalGraphQL :
func (version *Version) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*version = Version(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON :
func (version Version) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(version)), nil
}
