package ether

import (
	"fmt"
	"testing"
)

//TestDefaultAccountAddress 测试账户地址
var (
	TestDefaultAccountAddress    = "0x915dDe6e87dEF17FB1FE09A7710FC65d4b855a08"
	TestDefaultAccountPrivateKey = "8848da9feb2008022e691185e43d8603adced3bcf24a976901db44ddc4521d4d"
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
		"0xf7CDB8BDcfe20498eA60de0c7ddd933f01461409",
		1000000000000000000,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
