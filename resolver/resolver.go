package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/onsi/ginkgo"

	// Me cago en san blas
	_ "github.com/onsi/gomega"
)

// Resolver : root resolver
type Resolver struct {
	db *gorm.DB
}

// NewRoot : generates query root entry point
func NewRoot(db *gorm.DB) *Resolver {
	var _ = ginkgo.Describe("Resolver", func() {

	})

	return &Resolver{db: db}
}

// getOptionalStrValue : helper that check if string if a pointer or not
func getOptionalStrValue(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

// getOptionalBoolValue : helper that check if bool if a pointer or not
func getOptionalBoolValue(val *bool) bool {
	if val == nil {
		return false
	}
	return *val
}
