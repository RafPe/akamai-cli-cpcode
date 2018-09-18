package main

import (
	"fmt"

	edgegrid "github.com/RafPe/go-edgegrid"
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

	outputTableContracts(contracts)

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

	if output == "json" {
		OutputJSON(groups)
	} else {
		outputTableGroups(groups)
	}

	return nil
}

/*
	listProducts
*/
func cmdListProducts(c *cli.Context) error {
	return listProducts(c)
}

func listProducts(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")

	products, _, err := apiClient.PropertyAPI.ListPropertyAPIProducts(contractID)
	if err != nil {
		fmt.Println(err)
	}

	if output == "json" {
		OutputJSON(products)
	} else {
		outputTableProducts(products)
	}

	return nil
}

/*
	listCPcodes
*/
func cmdListCPcodes(c *cli.Context) error {
	return listCPcodes(c)
}

func listCPcodes(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")
	verifyArgumentByName(c, "groupID")

	cpcodes, _, err := apiClient.PropertyAPI.ListPropertyAPICPCodes(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	if output == "json" {
		OutputJSON(cpcodes)
	} else {
		outputTableCPCodes(cpcodes)
	}

	return nil
}

/*
	createCPcode
*/

func cmdCreateCPcode(c *cli.Context) error {
	return createCPcode(c)
}

func createCPcode(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")
	verifyArgumentByName(c, "groupID")
	verifyArgumentByName(c, "CPcodeName")
	verifyArgumentByName(c, "ProductID")

	newCPcode := &edgegrid.PropertyAPICPCodeNew{
		CpcodeName: CPcodeName,
		ProductID:  product,
	}

	_, err := apiClient.PropertyAPI.NewPropertyAPICPcode(newCPcode, contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ok")

	return nil
}
