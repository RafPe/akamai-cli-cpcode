package main

import (
	"fmt"
	"log"
	"os"

	edgegrid "github.com/RafPe/go-edgegrid"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/urfave/cli"
)

var (
	apiClient                                *edgegrid.Client
	apiClientOpts                            *edgegrid.ClientOptions
	homeDir, output, version                 string
	groupID, contractID, CPcodeName, product string
)

func main() {

	/*
		Sets default value for credentials configuration file
		to be pointing to ~/.edgerc
	*/
	homeDir, _ = homedir.Dir()
	homeDir += string(os.PathSeparator) + ".edgerc"

	/*
		Initialize values with using ENV variables either defaults
		AKAMAI_EDGERC_CONFIG  : for config file path
		AKAMAI_EDGERC_SECTION : for section
	*/
	apiClientOpts := &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = getEnv(string(edgegrid.EnvVarEdgercPath), homeDir)
	apiClientOpts.ConfigSection = getEnv(string(edgegrid.EnvVarEdgercSection), "default")

	/*
		Sets default values for app and global flags
	*/
	appName := "akamai-cpcodes"

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai CP codes"
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &apiClientOpts.ConfigSection,
			EnvVar:      string(edgegrid.EnvVarEdgercSection),
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       homeDir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &apiClientOpts.ConfigPath,
			EnvVar:      string(edgegrid.EnvVarEdgercPath),
		},
		cli.StringFlag{
			Name:        "output",
			Value:       "table",
			Usage:       "Defines output type ( json | table ) ",
			Destination: &output,
		},
	}

	app.Before = func(c *cli.Context) error {

		// create new Akamai API client
		apiClient = edgegrid.NewClient(nil, apiClientOpts)

		// if err != nil {
		// 	return cli.NewExitError(errorProfile, 1)
		// }

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Creates new cpcode",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "contractID",
					Usage:       "",
					Destination: &contractID,
				},
				cli.StringFlag{
					Name:        "groupID",
					Usage:       "",
					Destination: &groupID,
				},
				cli.StringFlag{
					Name:        "CPcodeName",
					Usage:       "",
					Destination: &CPcodeName,
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
							Name:        "contractID",
							Usage:       "",
							Destination: &contractID,
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
							Name:        "contractID",
							Usage:       "",
							Destination: &contractID,
						},
						cli.StringFlag{
							Name:        "groupID",
							Usage:       "",
							Destination: &groupID,
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

func verifyArgumentByName(c *cli.Context, argName string) {
	if c.String(argName) == "" {
		log.Fatal(fmt.Sprintf("Please provide required argument(s)! [ %s ]", argName))
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
