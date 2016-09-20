package opsman_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpsman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opsman Suite")
}
