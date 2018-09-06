package resolver_test

import (
	"changelog/config"
	"changelog/context"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *gorm.DB
var xx string

//This function will be called by the go test or ginkgo commands
func TestResolver(t *testing.T) {
	xx = "test resolver. "
	//Ginkgo signals failure
	RegisterFailHandler(Fail)
	// Starts the test suite. Will automatically fail testing.T if any of the specs fail
	RunSpecs(t, "Resolver Suite")
}

var _ = BeforeSuite(func() {
	c, err1 := config.LoadConfig("./../config")
	Expect(err1).NotTo(HaveOccurred())

	db, err2 := context.OpenDB(c)
	Expect(err2).NotTo(HaveOccurred())

	err3 := db.DB().Ping()
	Expect(err3).NotTo(HaveOccurred())

	xx += "before suite."

	log.Printf("XXXXXXXXXXXXXXXXXXXXXX BEFORES SUITE. DB = %v, xx = %v", db, xx)
})
var _ = AfterSuite(func() {
	xx += "after suite."
	db.Close()
	log.Printf("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC AFTER SUITE: DB = %v, xx = %v", db, xx)
})
