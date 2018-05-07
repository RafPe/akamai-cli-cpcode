package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	edgegrid "github.com/RafPe/go-edgegrid"
)

func outputTableCPCodes(cpcodes *edgegrid.PropertyAPICPCodes) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t ProductIDs \t Created \t Name"))
	for _, singleCPcoode := range cpcodes.Cpcodes.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s \t %s \t %s", singleCPcoode.CpcodeID, singleCPcoode.ProductIds, singleCPcoode.CreatedDate, singleCPcoode.CpcodeName))
	}

	w.Flush()

}

func outputTableProducts(products *edgegrid.PropertyAPIProducts) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t Name"))
	for _, singleProduct := range products.Products.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s", singleProduct.ProductID, singleProduct.ProductName))
	}

	w.Flush()

}

func outputTableGroups(groups *edgegrid.PropertyAPIGroups) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t ContractID \t Name"))
	for _, singleGroup := range groups.Groups.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s \t %s", singleGroup.GroupID, singleGroup.ContractIds, singleGroup.GroupName))
	}

	w.Flush()

}

func outputTableContracts(contracts *edgegrid.PropertyAPIContracts) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t Name"))
	for _, singleContract := range contracts.Contracts.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s", singleContract.ContractID, singleContract.ContractTypeName))
	}

	w.Flush()

}

// OutputJSON displays output of query for alerts in JSON format
//
// output
func OutputJSON(input interface{}) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
