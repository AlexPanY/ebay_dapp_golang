package main

import (
	"fmt"
	"log"

	"ebay_dapp_golang/ether"
)

func main() {
	account := ether.NewAccount(4, "0x02dAfB3c330715A45B1C44575397cdcD36ccf30d")
	b, err := account.GetBalanceByAddress()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(fmt.Sprintf("Account Balance=%v", b))
}
