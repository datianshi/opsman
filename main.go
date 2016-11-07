package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	const VERSION string = "0.0.7"
	app := &cli.App{
		Name:     "opsman-cli",
		HelpName: "opsman-cli",
		Version:  VERSION,
	}
	app.Commands = []*cli.Command{
		{
			Name:   "token",
			Usage:  "retieve token",
			Action: RetrieveToken,
			Flags:  []cli.Flag{OpsManagerURLFlag, Username, Password, SkipSSL},
		},
		{
			Name:        "pivnet",
			Usage:       "pivnet",
			Subcommands: []*cli.Command{
				{
					Name:   "download",
					Usage:  "download product",
					Action: DownloadProduct,
					Flags:  []cli.Flag{ProductURL, PivnetToken, SaveProductTo},
				},
				{
					Name:   "latest-release",
					Usage:  "Latest Release",
					Action: LatestProduct,
					Flags:  []cli.Flag{ProductNAME, PivnetToken},
				},
			},
		},
		{
			Name:  "upload",
			Usage: "upload",
			Subcommands: []*cli.Command{
				{
					Name:   "product",
					Usage:  "product",
					Action: UploadProduct,
					Flags:  []cli.Flag{OpsManagerURLFlag, Username, Password, UploadProductFrom, SkipSSL},
				},
				{
					Name:   "stemcell",
					Usage:  "stemcell",
					Action: UploadStemcell,
					Flags:  []cli.Flag{OpsManagerURLFlag, Username, Password, UploadStemcellFrom, SkipSSL},
				},
			},
		},
	}
	app.Run(os.Args)
}
