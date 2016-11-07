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

const TOKEN string = "hsbXaZnyquqaG9Nuwizb"

const LATEST_RELEASE_RESP=`
{
   "id":2799,
   "version":"1.8.13",
   "release_type":"Security Release",
   "release_date":"2016-11-02",
   "release_notes_url":"http://docs.pivotal.io/pivotalcf/1-8/pcf-release-notes/runtime-rn.html",
   "availability":"All Users",
   "description":"Please refer to the release notes",
   "eula":{
      "id":68,
      "slug":"pivotal_software_eula",
      "name":"Pivotal Software EULA",
      "_links":{
         "self":{
            "href":"https://network.pivotal.io/api/v2/eulas/68"
         }
      }
   },
   "end_of_support_date":"2017-06-30",
   "eccn":"5D002",
   "license_exception":"ENC",
   "controlled":true,
   "product_files":[
      {
         "id":6741,
         "aws_object_key":"product_files/Pivotal-CF/open_source_license_PCF-Elastic-Runtime-1.8.txt",
         "file_version":"1.0",
         "name":"PCF Elastic Runtime 1.8 License",
         "_links":{
            "self":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/6741"
            },
            "download":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/6741/download"
            },
            "signature_file_download":{
               "href":null
            }
         }
      },
      {
         "id":8517,
         "aws_object_key":"product_files/Pivotal-CF/DiegoWindows-1.8.3.zip",
         "file_version":"1.8.3",
         "name":"DiegoWindows",
         "_links":{
            "self":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8517"
            },
            "download":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8517/download"
            },
            "signature_file_download":{
               "href":null
            }
         }
      },
      {
         "id":8852,
         "aws_object_key":"product_files/Pivotal-CF/pcf_1.8.13-build.3_cloudformation.json",
         "file_version":"1.8.13-build.3",
         "name":"PCF Cloudformation for AWS Setup",
         "_links":{
            "self":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8852"
            },
            "download":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8852/download"
            },
            "signature_file_download":{
               "href":null
            }
         }
      },
      {
         "id":8851,
         "aws_object_key":"product_files/Pivotal-CF/cf-1.8.13-build.3.pivotal",
         "file_version":"1.8.13-build.3",
         "name":"PCF Elastic Runtime",
         "_links":{
            "self":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8851"
            },
            "download":{
               "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8851/download"
            },
            "signature_file_download":{
               "href":null
            }
         }
      }
   ],
   "file_groups":[

   ],
   "updated_at":"2016-11-04T01:26:01.626Z",
   "_links":{
      "self":{
         "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799"
      },
      "eula_acceptance":{
         "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/eula_acceptance"
      },
      "product_files":{
         "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files"
      },
      "file_groups":{
         "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/file_groups"
      },
      "user_groups":{
         "href":"https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/user_groups"
      }
   }
}
`


var _ = Describe("Pivnet", func() {
	var _ = Describe("Given a Pivnet product", func() {

		var pivnet *Pivnet
		var server *ghttp.Server
		var statusCode int
		var write_content *[]byte

		Context("When Pivnet request the product url", func() {
			var DOWNLOAD_PATH string = "/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download"
			var dest *bytes.Buffer
			BeforeEach(func() {
				statusCode = 200
				server = ghttp.NewServer()
				dest = &bytes.Buffer{}
				b := []byte("Hello_World")
				write_content = &b
				pivnet = &Pivnet{
					Token: TOKEN,
				}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", DOWNLOAD_PATH),
					ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Token %s", TOKEN)),
					ghttp.RespondWithPtr(&statusCode, write_content),
				))
			})

			It("Should download the product successfully", func() {
				err := pivnet.Download(dest, fmt.Sprintf("%s%s", server.URL(), DOWNLOAD_PATH))
				Ω(err).Should(BeNil())
				Ω(opsmantest.CompareMd5(dest, write_content)).Should(BeTrue())

			})
			AfterEach(func() {
				server.Close()
			})
		})

		Context("When Pivnet request the product latest version", func() {
			var productName string = "elastic-runtime"
			BeforeEach(func() {
				statusCode = 200
				server = ghttp.NewServer()
				b := []byte(LATEST_RELEASE_RESP)
				write_content = &b
				pivnet = &Pivnet{
					Token: TOKEN,
					PivURL: server.URL(),
				}
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/api/v2/products/elastic-runtime/releases/latest"),
					ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Token %s", TOKEN)),
					ghttp.RespondWithPtr(&statusCode, write_content),
				))
			})

			It("Should return the latest product", func() {
				product, err := pivnet.LatestProduct(productName)
				Ω(err).Should(BeNil())
				Ω(product.Id).Should(Equal(int64(2799)))
				Ω(product.AcceptUrl).Should(Equal("https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/eula_acceptance"))
				Ω(product.Files[0].Name).Should(Equal("PCF Elastic Runtime 1.8 License"))
				Ω(product.Files[0].DownloadUrl).Should(Equal("https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/6741/download"))
				Ω(product.Files[3].Name).Should(Equal("PCF Elastic Runtime"))
				Ω(product.Files[3].DownloadUrl).Should(Equal("https://network.pivotal.io/api/v2/products/elastic-runtime/releases/2799/product_files/8851/download"))
			})
			AfterEach(func() {
				server.Close()
			})
		})

	})
})
