package main

import (
	"os"
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"github.com/datianshi/opsman/uaa"
	"github.com/datianshi/opsman/pivnet"
	"github.com/cloudfoundry/gofileutils/fileutils"
)

func main() {
	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name: "token",
			Usage: "retieve token",
			Action: RetrieveToken,
			Flags: []cli.Flag{OpsManagerURLFlag, Username, Password, SkipSSL},
		},
		{
			Name: "download",
			Usage: "download product",
			Action: DownloadProduct,
			Flags: []cli.Flag{ProductURL, PivnetToken, SaveProductTo},
		},
	}
	app.Run(os.Args)
}

var OpsManagerURLFlag *cli.StringFlag = &cli.StringFlag{
	Name: "opsmanurl",
	Aliases: []string{"ops"},
	Usage: "ops manager url",
	EnvVars: []string{"OPS_MAN_URL"},
}

var ProductURL *cli.StringFlag = &cli.StringFlag{
	Name: "producturl",
	Aliases: []string{"prod"},
	Usage: "pivnet product url",
	EnvVars: []string{"PRODUCT_URL"},
}

var PivnetToken *cli.StringFlag = &cli.StringFlag{
	Name: "token",
	Aliases: []string{"t"},
	Usage: "pivnet token",
	EnvVars: []string{"PIVNET_TOKEN"},
}

var Username *cli.StringFlag = &cli.StringFlag{
	Name: "username",
	Aliases: []string{"u"},
	Usage: "ops manager username",
	EnvVars: []string{"OPS_MAN_USERNAME"},
}

var Password *cli.StringFlag = &cli.StringFlag{
	Name: "password",
	Aliases: []string{"p"},
	Usage: "ops manager username",
	EnvVars: []string{"OPS_MAN_PASSWORD"},
}

var SkipSSL *cli.BoolFlag = &cli.BoolFlag{
	Name: "skipssl",
	Usage: "skipssl",
	EnvVars: []string{"SKIP_SSL"},
}

var SaveProductTo *cli.StringFlag = &cli.StringFlag{
	Name: "dest",
	Aliases: []string{"d"},
	Usage: "Save the product file to",
	EnvVars: []string{"SAVE_PRODUCT_TO"},
}

var UploadProductFrom *cli.StringFlag = &cli.StringFlag{
	Name: "from",
	Usage: "Upload product from",
	Aliases: []string{"f"},
	EnvVars: []string{"UPLOAD_PRODUCT_FROM"},
}

func DownloadProduct(c *cli.Context) (err error){
	productURL := c.String("producturl")
	token := c.String("token")
	saveProductTo:= c.String("dest")
	file, err:= fileutils.Open(saveProductTo)
	defer file.Close()
	if(err!=nil){
		return
	}
	pivnet := pivnet.Pivnet{
		URL: productURL,
		Token: token,
	}
	err = pivnet.Download(file)
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
		URL: fmt.Sprintf("%s/uaa", opsManagerURL),
		SkipSsl: skipSsl,
	}
	token, err := uaa.GetToken();
	if (err != nil) {
		return err
	}
	fmt.Println(token)
	return nil
}
