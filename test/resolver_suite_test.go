package resolver_test

import (
	"changelog/config"
	"changelog/context"
	"testing"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *gorm.DB

func TestResolver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resolver Suite")
}

var _ = BeforeSuite(func() {
	c, err1 := config.LoadConfig("./../config")
	Expect(err1).NotTo(HaveOccurred())

	db, err2 := context.OpenDB(c)
	Expect(err2).NotTo(HaveOccurred())

	err3 := db.DB().Ping()
	Expect(err3).NotTo(HaveOccurred())

})
