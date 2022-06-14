package ether

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	AccountID uint64
	Address   string
}

//NewAccount
func NewAccount(accountID uint64, address string) Account {
	return Account{
		AccountID: accountID,
		Address:   address,
	}
}

//GetAccountByID
func GetAccountByID(accountID uint64) Account {
	return Account{
		AccountID: accountID,
		Address:   "0x71c7656ec7ab88b098defb751b7401b5f6d8976f",
	}
}

//GetBalanceByAddress
func (a *Account) GetBalanceByAddress() (*big.Int, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	accountAddress := common.HexToAddress(a.Address)

	balance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
