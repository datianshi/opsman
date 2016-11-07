package main

import (
	"fmt"
	"os"

	"github.com/datianshi/opsman/opsman"
	"github.com/datianshi/opsman/pivnet"
	"github.com/datianshi/opsman/uaa"
	"gopkg.in/urfave/cli.v2"
	"encoding/json"
)

var OpsManagerURLFlag *cli.StringFlag = &cli.StringFlag{
	Name:    "opsmanurl",
	Aliases: []string{"ops"},
	Usage:   "ops manager url",
	EnvVars: []string{"OPS_MAN_URL"},
}

var ProductURL *cli.StringFlag = &cli.StringFlag{
	Name:    "producturl",
	Aliases: []string{"prod"},
	Usage:   "pivnet product url",
	EnvVars: []string{"PRODUCT_URL"},
}

var ProductNAME *cli.StringFlag = &cli.StringFlag{
	Name:    "productname",
	Aliases: []string{"name"},
	Usage:   "pivnet product name",
	EnvVars: []string{"PRODUCT_NAME"},
}

var PivnetToken *cli.StringFlag = &cli.StringFlag{
	Name:    "token",
	Aliases: []string{"t"},
	Usage:   "pivnet token",
	EnvVars: []string{"PIVNET_TOKEN"},
}

var Username *cli.StringFlag = &cli.StringFlag{
	Name:    "username",
	Aliases: []string{"u"},
	Usage:   "ops manager username",
	EnvVars: []string{"OPS_MAN_USERNAME"},
}

var Password *cli.StringFlag = &cli.StringFlag{
	Name:    "password",
	Aliases: []string{"p"},
	Usage:   "ops manager username",
	EnvVars: []string{"OPS_MAN_PASSWORD"},
}

var SkipSSL *cli.BoolFlag = &cli.BoolFlag{
	Name:    "skipssl",
	Usage:   "skipssl",
	EnvVars: []string{"SKIP_SSL"},
}

var SaveProductTo *cli.StringFlag = &cli.StringFlag{
	Name:    "dest",
	Aliases: []string{"d"},
	Usage:   "Save the product file to",
	EnvVars: []string{"SAVE_PRODUCT_TO"},
}

var UploadProductFrom *cli.StringFlag = &cli.StringFlag{
	Name:    "from",
	Usage:   "Upload product from",
	Aliases: []string{"f"},
	EnvVars: []string{"UPLOAD_PRODUCT_FROM"},
}

var UploadStemcellFrom *cli.StringFlag = &cli.StringFlag{
	Name:    "from",
	Usage:   "Upload stemcell from",
	Aliases: []string{"f"},
	EnvVars: []string{"UPLOAD_STEMCELL_FROM"},
}

func DownloadProduct(c *cli.Context) (err error) {
	productURL := c.String("producturl")
	token := c.String("token")
	saveProductTo := c.String("dest")
	file, err := os.Create(saveProductTo)
	defer file.Close()
	if err != nil {
		return
	}
	pivnet := pivnet.Pivnet{
		PivURL: "https://network.pivotal.io/",
		Token: token,
	}
	err = pivnet.Download(file, productURL)
	return
}

func LatestProduct(c *cli.Context) (err error) {
	productName := c.String("productname")
	token := c.String("token")
	if err != nil {
		return
	}
	pivnet := pivnet.Pivnet{
		PivURL: "https://network.pivotal.io/",
		Token: token,
	}
	product, err := pivnet.LatestProduct(productName)
	if(err!=nil){
		return
	}
	b, err := json.Marshal(product)
	if(err!=nil){
		return
	}
	fmt.Println(string(b))
	return
}

func RetrieveToken(c *cli.Context) error {
	opsManagerURL := c.String("opsmanurl")
	username := c.String("username")
	password := c.String("password")
	skipSsl := c.Bool("skipssl")

	uaa := &uaa.UAA{
		Username: username,
		Password: password,
		URL:      opsManagerURL,
		SkipSsl:  skipSsl,
	}
	token, err := uaa.GetToken()
	if err != nil {
		return err
	}
	fmt.Println(token)
	return nil
}

type UploadMethod func(*opsman.OpsMan, *os.File) error

func upload(c *cli.Context, method UploadMethod) (err error) {
	opsManagerURL := c.String("opsmanurl")
	username := c.String("username")
	password := c.String("password")
	uploadProductFrom := c.String("from")
	skipSsl := c.Bool("skipssl")

	uaa := &uaa.UAA{
		Username: username,
		Password: password,
		URL:      opsManagerURL,
		SkipSsl:  skipSsl,
	}

	opsMan := opsman.CreateOpsman(opsManagerURL, skipSsl, uaa)
	file, err := os.Open(uploadProductFrom)
	defer file.Close()
	if err != nil {
		return
	}
	return method(opsMan, file)
}

func UploadProduct(c *cli.Context) (err error) {
	var method UploadMethod = func(opsman *opsman.OpsMan, file *os.File) error {
		return opsman.UploadProduct(file)
	}
	return upload(c, method)
}

func UploadStemcell(c *cli.Context) (err error) {
	var method UploadMethod = func(opsman *opsman.OpsMan, file *os.File) error {
		return opsman.UploadStemcell(file)
	}
	return upload(c, method)
}
