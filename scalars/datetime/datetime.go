package datetime

import (
	"fmt"
	"strconv"
)

// DateTime :
type DateTime string

// ImplementsGraphQLType :
func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// UnmarshalGraphQL :
func (datetime *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*datetime = DateTime(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON :
func (datetime DateTime) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(datetime)), nil
}
