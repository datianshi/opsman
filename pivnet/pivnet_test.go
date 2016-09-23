package pivnet_test

import (
	. "github.com/datianshi/opsman/pivnet"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Given a Pivnet product and token", func() {
	var dest *os.File
	var pivnet *Pivnet
	BeforeEach(func() {
		//dest, _= ioutil.TempFile("", "test")
		dest, _ = os.Create("/tmp/file")
		pivnet = &Pivnet{
			URL:   "https://network.pivotal.io/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download",
			Token: "hsbXaZnyquqaG9Nuwizb",
		}
	})
	Context("When Pivnet request the product", func() {
		It("Should download the product successfully", func() {
			defer dest.Close()
			err := pivnet.Download(dest)
			Ω(err).Should(BeNil())
		})
	})
	//AfterEach(func() {
	//	os.Remove(dest.Name())
	//})
})
