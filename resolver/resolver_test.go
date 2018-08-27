package resolver_test

import (
	graphql "github.com/graph-gophers/graphql-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"changelog/model"
	. "changelog/resolver"
)

// The var _= ... trick allows to evaluate the Describe at the top level
// without having to wrap it in a  func init() {}
var _ = Describe("Resolver", func() {

	var (
		// Variable declaration
		data      model.UpdateChangeInput
		cResolver *Resolver
		result    *ChangeResolver
	)
	type inputData struct {
		Input *model.UpdateChangeInput
	}

	BeforeEach(func() {
		// TODO:
		// Set up state for specs.
	})
	It("data", func() {
		// Set up a single spec.
		data = model.UpdateChangeInput{
			Code:    "ch2",
			Message: "messageTest",
			Project: "projectTest",
			Type:    model.ChangeType("FIXED"),
		}
	})

	// Example

	Describe("Updating an existing change with new data", func() {
		Context("Modifing data from a manually created change", func() {
			It("should have the same values in DB after the UpdateChange function", func() {
				//UpdateChange
				s := inputData{Input: &data}
				result = cResolver.UpdateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(data.Code)))
			})
		})
	})
})
