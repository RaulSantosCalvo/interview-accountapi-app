package main

import (
	"fmt"
	"log"

	"github.com/RaulSantosCalvo/interview-accountapi/accountapi"
)

const accountId string = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

var account string = fmt.Sprintf(`{"attributes":{"bank_id":"123456","bank_id_code":"GBDSC","base_currency":"GBP","bic":"EXMPLGB2XXX","country":"GB","name":["test_account"]},"id":"%s","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts"}`, accountId)

func create() {
	fmt.Println("create")
	result, err := accountapi.Create(account)
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
