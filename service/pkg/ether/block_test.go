package ether

import (
	"fmt"
	"testing"
)

//TestGetHeaderByNumber
//NOTE: go test -run TestGetHeaderByNumber -v
func TestGetHeaderByNumber(t *testing.T) {
	header, err := GetHeaderByNumber(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("Header number = %+v", header.Number.String()))
}
