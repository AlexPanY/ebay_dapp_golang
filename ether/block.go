package ether

import (
	"context"
	"math/big"

	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//GetHeaderByNumber 获取
func GetHeaderByNumber(number *big.Int) (*types.Header, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	header, err := client.HeaderByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	return header, nil
}
