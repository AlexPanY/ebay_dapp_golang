package ether

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	defaultRawURL = "http://127.0.0.1:8545"
)

//NewEtherClientConn
func NewEtherClientConn() (*ethclient.Client, error) {
	client, err := ethclient.Dial(defaultRawURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}
