package resolver_test

import (
	graphql "github.com/graph-gophers/graphql-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"changelog/config"
	"changelog/context"
	"changelog/model"
	"changelog/resolver"
)

// The var _= ... trick allows to evaluate the Describe at the top level
// without having to wrap it in a  func init() {}
var _ = Describe("Resolver", func() {
	var (
		// Variable declaration
		UpdateData model.UpdateChangeInput
		CreateData model.CreateChangeInput
		cResolver  resolver.Resolver
	)

	BeforeEach(func() {
		// Load Config & DB
		c, err1 := config.LoadConfig("./../config")
		Expect(err1).NotTo(HaveOccurred())

		db, err2 := context.OpenDB(c)
		Expect(err2).NotTo(HaveOccurred())

		err3 := db.DB().Ping()
		Expect(err3).NotTo(HaveOccurred())

		// Resolver.resolver init
		cResolver = *resolver.NewRoot(db)

		// TODO:
		// Set up state for specs.
		UpdateData = model.UpdateChangeInput{
			Code:    "ch2",
			Message: "messageTest",
			Project: "projectTest",
			Type:    model.ChangeType("CHANGED"),
		}

		newCode := "ch7"
		CreateData = model.CreateChangeInput{
			Code:    &newCode,
			Message: "newMessage",
			Project: "newProject",
			Type:    model.ChangeType("ADDED"),
		}
	})

	It("data", func() {
		// Set up a singel spec
	})

	// Example

	Describe("Updating an existing change with new data. ", func() {
		Context("Modifing data from a manually created change", func() {
			It("should have the same values in DB after the UpdateChange function", func() {
				//UpdateChange
				s := struct{ Input *model.UpdateChangeInput }{&UpdateData}
				result := *cResolver.UpdateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(UpdateData.Code)))
			})
		})
	})

	Describe("Creating a new change", func() {
		Context("adding data from a manually created node", func() {
			It("should have the following values: Code: ch8\n Message: messageTest\n Project: projectTest\n Type: FIXED after retrieving it form the DB", func() {
				s := struct{ Input *model.CreateChangeInput }{&CreateData}
				result := *cResolver.CreateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(*CreateData.Code)))
			})
		})
	})
})
