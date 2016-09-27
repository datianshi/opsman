package pivnet_test

import (
	. "github.com/datianshi/opsman/pivnet"

	"bytes"
	"fmt"

	"github.com/datianshi/opsman/opsmantest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

const DOWNLOAD_PATH string = "/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download"
const TOKEN string = "hsbXaZnyquqaG9Nuwizb"

var _ = Describe("Given a Pivnet product and token", func() {
	var dest *bytes.Buffer
	var pivnet *Pivnet
	var server *ghttp.Server
	var statusCode int
	var write_content *[]byte

	BeforeEach(func() {
		statusCode = 200
		server = ghttp.NewServer()
		dest = &bytes.Buffer{}
		b := []byte("Hello_World")
		write_content = &b
		pivnet = &Pivnet{
			URL:   fmt.Sprintf("%s%s", server.URL(), DOWNLOAD_PATH),
			Token: TOKEN,
		}
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", DOWNLOAD_PATH),
			ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Token %s", TOKEN)),
			ghttp.RespondWithPtr(&statusCode, write_content),
		))
	})

	Context("When Pivnet request the product", func() {
		It("Should download the product successfully", func() {
			err := pivnet.Download(dest)
			Ω(err).Should(BeNil())
			Ω(opsmantest.CompareMd5(dest, write_content)).Should(BeTrue())

		})
	})

	AfterEach(func() {
		server.Close()
	})
})
