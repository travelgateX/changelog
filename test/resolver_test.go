package resolver_test

import (
	"changelog/config"
	"changelog/context"
	"changelog/model"
	"changelog/resolver"

	graphql "github.com/graph-gophers/graphql-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// The var _= ... trick allows to evaluate the Describe at the top level
// without having to wrap it in a  func init() {}
var _ = Describe("Resolver", func() {
	var (
		// Variable declaration
		UpdateData     model.UpdateChangeInput
		CreateData     model.CreateChangeInput
		DeleteData     model.DeleteChangeInput
		cResolver      resolver.Resolver
		CreateDataTest model.CreateChangeInput
		UpdateDataTest model.UpdateChangeInput
		DeleteDataTest model.DeleteChangeInput
	)

	BeforeEach(func() {
		// Load Config & DB
		// TODO:
		// Set up state for specs.
		c, err1 := config.LoadConfig("./../config")
		Expect(err1).NotTo(HaveOccurred())

		db, err2 := context.OpenDB(c)
		Expect(err2).NotTo(HaveOccurred())
		err3 := db.DB().Ping()
		Expect(err3).NotTo(HaveOccurred())

		cResolver = *resolver.NewRoot(db)

		newCode := "NewCode1"
		CreateData = model.CreateChangeInput{
			Code:    &newCode,
			Message: "newMessage",
			Project: "newProject",
			Type:    model.ChangeType("ADDED"),
		}
		UpdateData = model.UpdateChangeInput{
			Code:    newCode,
			Message: "UpdatedMessageTest",
			Project: "UpdatedProjectTest",
			Type:    model.ChangeType("CHANGED"),
		}
		DeleteData = model.DeleteChangeInput{
			Code: newCode,
		}
		newCodeTest := "NewCode2"
		CreateDataTest = model.CreateChangeInput{
			Code:    &newCodeTest,
			Message: "NewMessage",
			Project: "NewProject",
			Type:    model.ChangeType("ADDED"),
		}

		UpdateDataTest = model.UpdateChangeInput{
			Code:    newCodeTest,
			Message: "Updated message",
			Project: "Updated project",
			Type:    model.ChangeType("CHANGED"),
		}

		DeleteDataTest = model.DeleteChangeInput{
			Code: newCodeTest,
		}

	})
	JustBeforeEach(func() {
	})

	// CreateChange function test
	Describe("Creating a new change", func() {
		Context("adding data from a manually created node", func() {
			It("should have the following values: Code: ch8\n Message: messageTest\n Project: projectTest\n Type: FIXED\n after retrieving it form the DB", func() {
				s := struct{ Input *model.CreateChangeInput }{&CreateData}
				result := *cResolver.CreateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(*CreateData.Code)))
			})
		})

		Context("if the data code is already stored in DB", func() {
			It("should return an advise message and the data in change should be null", func() {
				s := struct{ Input *model.CreateChangeInput }{&CreateData}
				result := *cResolver.CreateChange(s)
				type lf struct {
					Level *[]*model.AdviseMessageLevel
				}
				Expect(result.AdviseMessage(lf{nil}) != nil)

			})
		})
	})

	// UpdateChange function test
	Describe("Creating a new change input... ", func() {
		Context("if code is not already stored in DB", func() {
			It("should have the same values in DB after the UpdateChange function", func() {
				s := struct{ Input *model.UpdateChangeInput }{&UpdateData}
				result := *cResolver.UpdateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(UpdateData.Code)))
			})
		})

		Context("if trying to update non existing change", func() {
			It("should return an advise message and the data in change should be null", func() {
				// Use code to test error handling when code is duplicated
				wrongCode := model.UpdateChangeInput{
					Code: "asdas",
				}
				k := struct{ Input *model.UpdateChangeInput }{&wrongCode}
				result := *cResolver.UpdateChange(k)
				type lf struct {
					Level *[]*model.AdviseMessageLevel
				}
				Expect(result.AdviseMessage(lf{nil}) != nil)
			})
		})
	})

	// La función Unscoped no funciona debidamente
	Describe("Deleting record from DB", func() {
		Context("if code exists ", func() {
			It("should mark it as soft-deleted", func() {
				s := struct{ Input *model.DeleteChangeInput }{&DeleteData}
				result := *cResolver.DeleteChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(DeleteData.Code)))
			})
		})
	})
	Describe("Integration test", func() {
		Context("from pre-loaded data, perform create, update and delete functions", func() {
			It("should pass all the asserts", func() {

				c, err1 := config.LoadConfig("./../config")
				Expect(err1).NotTo(HaveOccurred())

				db, err2 := context.OpenDB(c)
				Expect(err2).NotTo(HaveOccurred())

				err3 := db.DB().Ping()
				Expect(err3).NotTo(HaveOccurred())

				cResolver = *resolver.NewRoot(db)

				// Add new data to DB
				s := struct{ Input *model.CreateChangeInput }{&CreateDataTest}
				result := *cResolver.CreateChange(s)
				_ = result
				// Fetch data
				var resultDB model.ChangeData
				st := *CreateDataTest.Code
				db.First(&resultDB, "code = ?", st)

				// Asert create
				Expect(resultDB.Code).To(Equal(graphql.ID(*CreateDataTest.Code)))
				Expect(resultDB.Message).To(Equal(CreateDataTest.Message))
				Expect(resultDB.Project).To(Equal(&CreateDataTest.Project))
				Expect(resultDB.Type).To(Equal(CreateDataTest.Type))

				// Update same data
				up := struct{ Input *model.UpdateChangeInput }{&UpdateDataTest}
				upRes := *cResolver.UpdateChange(up)
				_ = upRes
				// Fetch data
				db.First(&resultDB, "code = ?", UpdateDataTest.Code)
				// Assert Update
				Expect(resultDB.Code).To(Equal(graphql.ID(UpdateDataTest.Code)))
				Expect(resultDB.Message).To(Equal(UpdateDataTest.Message))
				Expect(resultDB.Project).To(Equal(&UpdateDataTest.Project))
				Expect(resultDB.Type).To(Equal(UpdateDataTest.Type))

				// Delete same data
				del := struct{ Input *model.DeleteChangeInput }{&DeleteDataTest}
				delRes := *cResolver.DeleteChange(del)
				_ = delRes
				// Fetch data
				db.Unscoped().Where("code = ?", DeleteDataTest.Code).First(&resultDB)
				// Assert Delete
				Expect(resultDB.DeletedAt != nil)
			})
		})
	})
})
