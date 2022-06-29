package ether

import (
	"fmt"
	"testing"
)

//TestDefaultAccountAddress 测试账户地址
var (
	TestDefaultAccountAddress    = "0x1BBDcBC03275b819Fb6E34Ad699be47213efB641"
	TestDefaultAccountPrivateKey = "41a7769df9a5b19951d91f96df8d43e1bccf871a432de3423a441591f2f3f19c"
)

//TestGetBalanceByAddress
//NOTE: go test -run TestGetBalanceByAddress -v
func TestGetBalanceByAddress(t *testing.T) {
	a := NewAccount(4, TestDefaultAccountAddress, TestDefaultAccountPrivateKey)
	balance, err := a.GetBalanceByAddress()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Account Balance=%v", balance))
}

//TestGetPendingBalanceByAddress
//NOTE: go test -run TestGetPendingBalanceByAddress -v
func TestGetPendingBalanceByAddress(t *testing.T) {
	a := NewAccount(4, TestDefaultAccountAddress, TestDefaultAccountPrivateKey)
	balance, err := a.GetPendingBalanceByAddress()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Account Pending Balance=%v", balance))
}

//TestTransferETHByAddress
//NOTE: go test -run TestTransferETHByAddress -v
func TestTransferETHByAddress(t *testing.T) {
	a := NewAccount(4, TestDefaultAccountAddress, TestDefaultAccountPrivateKey)
	err := a.TransferETHByAddress(
		"0x668396FB2Dcd72A9c9CF8A67EC06e7FfaDb49efB",
		1000000000000000000,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
