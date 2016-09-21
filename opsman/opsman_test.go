package opsman_test

import (
	. "github.com/datianshi/opsman/opsman"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = FDescribe("Given an OpsMan", func() {
	var opsMan *OpsMan
	var productFile *os.File

	BeforeEach(func() {
		opsMan = &OpsMan{
			Username: "admin",
			Password: "password",
			OpsManUrl: "https://opsmgr.haas-22.pez.pivotal.io",
			SkipSsl: true,
		}
		productFile, _ = os.Open("/tmp/apigee-cf-service-broker-0.2.1.pivotal")
	})

	Context("When upload a file", func() {
		It("Should not have any err happened", func() {
			err := opsMan.Upload(productFile)
			Ω(err).Should(BeNil())
		})
	})
	AfterEach(func(){
		productFile.Close()
	})
})
