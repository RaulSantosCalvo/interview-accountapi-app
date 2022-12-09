package main

import (
	"fmt"
	"log"

	"github.com/RaulSantosCalvo/interview-accountapi/accountapi"
)

const accountId string = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

// returns mocked account object
func getMockAccount(id string) accountapi.AccountData {
	country := "GB"
	account := accountapi.AccountData{
		ID:             id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &accountapi.AccountAttributes{
			Country:      &country,
			BaseCurrency: "GBP",
			BankID:       "123456",
			BankIDCode:   "GBDSC",
			Bic:          "EXMPLGB2XXX",
			Name:         []string{"Test", "Account"},
		},
	}
	return account
}

func create() {
	fmt.Println("create")
	result, err := accountapi.Create(getMockAccount(accountId))
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println("create result: ", result)
}

func fetch() {
	fmt.Println("fetch")
	result, err := accountapi.Fetch(accountId)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println("fetch result: ", result)
}

func delete() {
	fmt.Println("delete")
	result, err := accountapi.Delete(accountId, "0")
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println("delete result: ", result)
}

func main() {
	create()
	fetch()
	delete()
	fetch()
}
