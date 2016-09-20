package pivnet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPivnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pivnet Suite")
}
