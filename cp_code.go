package main

import (
	"fmt"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

/*
	listContracts
*/
func cmdListContracts(c *cli.Context) error {
	return listContracts(c)
}

func listContracts(c *cli.Context) error {
	// List all network lists
	contracts, _, err := apiClient.PropertyAPI.ListPropertyAPIContracts()
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(contracts)

	return nil
}

/*
	listGroups
*/
func cmdListGroups(c *cli.Context) error {
	return listGroups(c)
}

func listGroups(c *cli.Context) error {
	// List all network lists
	groups, _, err := apiClient.PropertyAPI.ListPropertyAPIGroups()
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(groups)

	return nil
}

/*
	listProducts
*/
func cmdListProducts(c *cli.Context) error {
	return listProducts(c)
}

func listProducts(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")

	products, _, err := apiClient.PropertyAPI.ListPropertyAPIProducts(contractID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(products)

	return nil
}

/*
	listCPcodes
*/
func cmdListCPcodes(c *cli.Context) error {
	return listCPcodes(c)
}

func listCPcodes(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")
	common.VerifyArgumentByName(c, "groupID")

	cpcodes, _, err := apiClient.PropertyAPI.ListPropertyAPICPCodes(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(cpcodes)

	return nil
}

/*
	createCPcode
*/

func cmdCreateCPcode(c *cli.Context) error {
	return createCPcode(c)
}

func createCPcode(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")
	common.VerifyArgumentByName(c, "groupID")
	common.VerifyArgumentByName(c, "CPcodeName")
	common.VerifyArgumentByName(c, "ProductID")

	newCPcode := &edgegrid.PropertyAPICPCodeNew{
		CpcodeName: CPcodeName,
		ProductID:  productID,
	}

	resp, err := apiClient.PropertyAPI.NewPropertyAPICPcode(newCPcode, contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(resp.Body)

	return nil
}
