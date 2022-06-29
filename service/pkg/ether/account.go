package ether

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	AccountID  uint64
	Address    string
	EthAddress common.Address
	PrivateKey string
}

//NewAccount
func NewAccount(accountID uint64, address, key string) Account {
	return Account{
		AccountID:  accountID,
		Address:    address,
		EthAddress: common.HexToAddress(address),
		PrivateKey: key,
	}
}

//GetAccountByID
func GetAccountByID(accountID uint64) Account {
	a := Account{
		AccountID: accountID,
		Address:   "0x71c7656ec7ab88b098defb751b7401b5f6d8976f",
	}
	a.EthAddress = common.HexToAddress(a.Address)
	return a
}

//TransferETHByAddress 通过ETH进行转账交易
func (a *Account) TransferETHByAddress(toAdderss string, amount int64) error {
	err := Transfer(a.PrivateKey, toAdderss, amount)
	if err != nil {
		return err
	}
	return nil
}

//GetBalanceByAddress
func (a *Account) GetBalanceByAddress() (*big.Int, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), a.EthAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

//GetPendingBalanceByAddress 获取账户待处理余额
func (a *Account) GetPendingBalanceByAddress() (*big.Int, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), a.EthAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
