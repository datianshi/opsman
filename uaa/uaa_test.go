package uaa_test

import (
	. "github.com/datianshi/opsman/uaa"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

const TOKEN_PATH = "/uaa/oauth/token"
const USERNAME = "admin"
const PASSWORD = "password"

var WRITE_CONTENT string = `{
	"access_token" : "helloworld"
}`

var _ = Describe("Uaa", func() {

	var server *ghttp.Server
	var uaaRequest *UAA
	var statusCode int

	BeforeEach(func() {
		statusCode = 200
		server = ghttp.NewServer()
		uaaRequest = &UAA{
			URL:      server.URL(),
			Username: USERNAME,
			Password: PASSWORD,
			SkipSsl:  true,
		}
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", TOKEN_PATH),
			ghttp.VerifyBasicAuth("opsman", ""),
			ghttp.RespondWithPtr(&statusCode, &WRITE_CONTENT),
		),)
	})

	Context("When server returned token successfully", func() {
		It("err should be nil", func() {
			_, err := uaaRequest.GetToken()
			Ω(err).Should(BeNil())
		})
		It("token should be returned", func() {
			token, _ := uaaRequest.GetToken()
			Ω(token).Should(Equal("helloworld"))
		})
	})


	AfterEach(func() {
		server.Close()
	})
})
