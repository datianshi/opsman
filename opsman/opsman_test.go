package opsman_test

import (
	. "github.com/datianshi/opsman/opsman"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"github.com/onsi/gomega/ghttp"
	"github.com/datianshi/opsman/opsmantest"
	"fmt"
	"github.com/datianshi/opsman/uaa"
)

const UPLOAD_PRODUCT string = "/uaa/oauth/token"
const TOKEN string = "Hello World"
const OPS_USERNAME = "admin"

var _ = Describe("Given an OpsMan", func() {
	var opsMan *OpsMan
	var statusCode int
	var productFile string
	var server *ghttp.Server
	var tokenIssuer uaa.TokenIssuer

	Describe("Upload a product", func() {
		BeforeEach(func() {
			server = ghttp.NewServer()
			tokenIssuer = &opsmantest.FakeTokenIssuer{
				Token: TOKEN,
				ErrorControl: nil,
			}
			statusCode = 200
			opsMan = CreateOpsman(server.URL(), false, tokenIssuer)
			productFile, _ = opsmantest.CreateGarbageFile("sdhfajhfasdasdhfajshdf")
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", UPLOAD_PRODUCT_PATH),
				ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Bearer %s", TOKEN)),
				opsmantest.VerifyUploadFile(productFile, "product[file]"),
				ghttp.RespondWithPtr(&statusCode, nil),
			))
		})

		Context("when the server respond with 200", func() {
			It("Should not have any err happened", func() {
				file, _ := os.Open(productFile)
				defer file.Close()
				err := opsMan.Upload(file)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when the server respond with 500", func() {
			BeforeEach(func() {
				statusCode = 500
			})
			It("Should not have  err happened", func() {
				file, _ := os.Open(productFile)
				defer file.Close()
				err := opsMan.Upload(file)
				Ω(err).Should(HaveOccurred())
			})
		})
		AfterEach(func() {
			server.Close()
			os.Remove(productFile)
		})
	})

})
