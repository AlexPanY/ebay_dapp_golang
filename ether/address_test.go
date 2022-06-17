package ether

import (
	"fmt"
	"testing"
)

//TestDefaultContractAddress 测试合约地址
var TestDefaultContractAddress = "0xE475B6c9d902621ff583c94ae66f12fb00644315"

//TestIsContractAddress
//NOTE: go test -run TestIsContractAddress -v
func TestIsContractAddress(t *testing.T) {
	err := IsContractAddress(TestDefaultContractAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
