package resolver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//This function will be called by the go test or ginkgo commands
func TestResolver(t *testing.T) {
	//Ginkgo signals failure
	RegisterFailHandler(Fail)
	// Starts the test suite. Will automatically fail testing.T if any of the specs fail
	RunSpecs(t, "Resolver Suite")
}
