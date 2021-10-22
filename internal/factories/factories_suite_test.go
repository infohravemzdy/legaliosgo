package factories_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testParams struct {
	testName   string
	testYear   int16
	testMonth  int16
	resultYear int16
}
type testScenario struct {
	title string
	tests []testParams
}

func TestFactories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factories Suite")
}
