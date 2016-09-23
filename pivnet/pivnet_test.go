package pivnet_test

import (
	. "github.com/datianshi/opsman/pivnet"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"io/ioutil"
)

var _ = FDescribe("Given a Pivnet product and token", func() {
	var dest *os.File
	var pivnet *Pivnet
	BeforeEach(func() {
		dest, _= ioutil.TempFile("", "test")
		pivnet = &Pivnet{
			URL: "https://network.pivotal.io/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download",
			Token: "wtKajG26LMsv3JDN_sCB",
		}
	})
	Context("When Pivnet request the product", func() {
		It("Should download the product successfully", func() {
			err:= pivnet.Download(dest)
			Ω(err).Should(BeNil())
		})
	})
	AfterEach(func(){
		os.Remove(dest.Name())
	})
})
