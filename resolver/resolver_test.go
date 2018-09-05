package resolver_test

import (
	"changelog/config"
	"changelog/context"
	"changelog/mocks"
	"changelog/model"
	"changelog/resolver"
	"log"

	"github.com/golang/mock/gomock"
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

		// Mock variables
		mockCtrl *gomock.Controller
		mockMut  *mocks.MockMutations
		consumer *consumer
	)

	BeforeEach(func() {
		// Load Config & DB

		mockCtrl = gomock.NewController(GinkgoT())
		mockMut = mocks.NewMockDBinterface
		// Set up state for specs.
		UpdateData = model.UpdateChangeInput{
			Code:    "ch3",
			Message: "messageTest",
			Project: "projectTest",
			Type:    model.ChangeType("CHANGED"),
		}

		newCode := "ch140"
		CreateData = model.CreateChangeInput{
			Code:    &newCode,
			Message: "newMessage",
			Project: "newProject",
			Type:    model.ChangeType("ADDED"),
		}

		DeleteData = model.DeleteChangeInput{
			Code: "ch7",
		}
		newCodeTest := "NewCode"
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
		// c, err1 := config.LoadConfig("./../config")
		// Expect(err1).NotTo(HaveOccurred())

		// db, err2 := context.OpenDB(c)
		// Expect(err2).NotTo(HaveOccurred())
		// err3 := db.DB().Ping()
		// Expect(err3).NotTo(HaveOccurred())

		xx += "justBeforeEach. "

		log.Printf("ZZZZZZZZZZZZZZZZZZZZZZZ JUST BEFORE EACH. DB = %v, xx = %v\n", db, xx)

		// Resolver.resolver init
		cResolver = *resolver.NewRoot(db)
	})

	// UpdateChange function test
	Describe("Creating a new change input... ", func() {
		Context("if code is not already stored in DB", func() {
			It("should have the same values in DB after the UpdateChange function", func() {
				//UpdateChange
				c, err1 := config.LoadConfig("./../config")
				Expect(err1).NotTo(HaveOccurred())

				xxxx, err2 := context.OpenDB(c)
				Expect(err2).NotTo(HaveOccurred())

				err3 := xxxx.DB().Ping()
				Expect(err3).NotTo(HaveOccurred())

				cResolver = *resolver.NewRoot(xxxx)

				xx += "Updatechange test! "
				log.Printf("XXXXXXXXXXXXXXXXXXXXXX JUST BEFORE EACH. DB = %v, xx = %v, c = %v\n", xxxx, xx, c)
				s := struct{ Input *model.UpdateChangeInput }{&UpdateData}
				result := *cResolver.UpdateChange(s)
				Expect(result.ChangeData().Code()).To(Equal(graphql.ID(UpdateData.Code)))
			})
		})

		Context("if trying to update non existing change", func() {
			It("should return an advise message and the data in change should be null", func() {

			})

		})
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
	// La función Unscoped no funciona debidamente
	Describe("Deleting record from DB", func() {
		Context("if code exists ", func() {
			It("should mark it as soft-deleted", func() {
				s := struct{ Input *model.DeleteChangeInput }{&DeleteData}
				result := *cResolver.DeleteChange(s)
				_ = result
				var z model.ChangeData
				db.Raw("SELECT * FROM changes WHERE code = ?", DeleteData.Code).Scan(&z)
			})
		})
	})
	Describe("Integration test", func() {
		Context("from pre-loaded data, perform create, update and delete functions", func() {
			It("should pass all the asserts", func() {

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
				Expect(resultDB.Project).To(Equal(CreateDataTest.Project))
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
				Expect(resultDB.Project).To(Equal(UpdateDataTest.Project))
				Expect(resultDB.Type).To(Equal(UpdateDataTest.Type))

				// Delete same data
				del := struct{ Input *model.DeleteChangeInput }{&DeleteDataTest}
				delRes := *cResolver.DeleteChange(del)
				_ = delRes
				// Fetch data
				db.Unscoped().Where("code=? ?", DeleteDataTest.Code).First(&resultDB)
				// Assert Delete
				Expect(resultDB.DeletedAt != nil)
			})
		})
	})
})
