package ether

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

//IsContractAddress 是否为合约地址
func IsContractAddress(address string) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if isVaild := re.MatchString(address); !isVaild {
		return errors.New("address is vaild.")
	}

	client, err := NewEtherClientConn()
	if err != nil {
		return err
	}

	ethAddress := common.HexToAddress(address)

	bytecode, err := client.CodeAt(context.Background(), ethAddress, nil)
	if err != nil {
		return err
	}

	fmt.Println(bytecode)
	return nil
}
