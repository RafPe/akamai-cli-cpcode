package main

import (
	"os"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var (
	apiClient       *edgegrid.Client
	appName, appVer string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to manage Akamai CPcodes", appVer)
	app.Flags = common.CreateFlags()

	app.Before = func(c *cli.Context) error {
		var err error

		// Provide struct details needed for apiClient init
		apiClientOpts := &edgegrid.ClientOptions{}
		apiClientOpts.ConfigPath = c.GlobalString("config")
		apiClientOpts.ConfigSection = c.GlobalString("section")
		apiClientOpts.DebugLevel = c.GlobalString("debug")
		apiClientOpts.AccountSwitchKey = c.GlobalString("ask")

		apiClient, err = common.EdgeClientInit(apiClientOpts)

		if err != nil {
			log.Fatalln(err)
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Creates new cpcode",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "contractID",
					Usage: "",
				},
				cli.StringFlag{
					Name:  "groupID",
					Usage: "",
				},
				cli.StringFlag{
					Name:  "ProductID",
					Usage: "",
				},
				cli.StringFlag{
					Name:  "CPcodeName",
					Usage: "",
				},
			},
			Action: cmdCreateCPcode,
		},
		{
			Name:  "list",
			Usage: "List cpcodes and account info",
			Subcommands: []cli.Command{
				{
					Name:     "contracts",
					Usage:    "List associated account contracts",
					Action:   cmdListContracts,
					Category: "Account actions",
				},
				{
					Name:     "groups",
					Usage:    "List associated account groups",
					Action:   cmdListGroups,
					Category: "Account actions",
				},
				{
					Name:  "products",
					Usage: "List associated contract products",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "contractID",
							Usage: "",
						},
					},
					Action:   cmdListProducts,
					Category: "Account actions",
				},
				{
					Name:  "cpcodes",
					Usage: "List associated contract/group cpcodes",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "contractID",
							Usage: "",
						},
						cli.StringFlag{
							Name:  "groupID",
							Usage: "",
						},
					},
					Action:   cmdListCPcodes,
					Category: "CPCodes actions",
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
