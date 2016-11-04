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

const TOKEN string = "Hello World"

var opsMan *OpsMan
var statusCode int
var productFile string
var server *ghttp.Server
var tokenIssuer uaa.TokenIssuer

type upload func(* os.File) error
var _ = Describe("Given an OpsMan", func() {

	var uploadProduct upload = func(file *os.File) error{
		return opsMan.UploadProduct(file)
	}

	var uploadStemcell upload = func(file *os.File) error{
		return opsMan.UploadStemcell(file)
	}

	describeUpload("Upload a product", UPLOAD_PRODUCT_PATH, UPLOAD_PRODUCT_FORM_PARAM, uploadProduct)
	describeUpload("Upload a stemcell", UPLOAD_STEMCELL_PATH, UPLOAD_STEMCELL_FORM_PARAM, uploadStemcell)
})

func describeUpload(describe, uploadUrl , uploadForm string, up upload){
	Describe(describe, func() {
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
				ghttp.VerifyRequest("POST", uploadUrl),
				ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Bearer %s", TOKEN)),
				opsmantest.VerifyUploadFile(productFile, uploadForm),
				ghttp.RespondWithPtr(&statusCode, nil),
			))
		})

		Context("when the server respond with 200", func() {
			It("Should not have any err happened", func() {
				file, _ := os.Open(productFile)
				defer file.Close()
				err := up(file)
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
				err := up(file)
				Ω(err).Should(HaveOccurred())
			})
		})
		AfterEach(func() {
			server.Close()
			os.Remove(productFile)
		})
	})

}
