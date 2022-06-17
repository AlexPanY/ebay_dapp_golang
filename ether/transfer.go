package ether

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

//Transfer
func Transfer(key, toAddress string, amount int64) error {

	client, err := NewEtherClientConn()
	if err != nil {
		return err
	}

	//1.解析私钥
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return err
	}

	//2.获取发送账户的公共地址
	//NOTE: Public Address 可以通过私钥派生出来
	publicKey := privateKey.Public()
	publicKeyECDSA, isOk := publicKey.(*ecdsa.PublicKey)
	if !isOk {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fmt.Sprintf("From Address = (%+v)", fromAddress))

	//3.获得帐户的随机数 `nonce`
	//NOTE: 每笔交易都需要一个 `nonce`,根据定义`nonce`仅使用一次
	//      `ethereum` 客户端提供了 PendingNonceAt() 方法，获取可用 `nonce` 值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	transferAmount := big.NewInt(amount)

	//4.设置交易的 `Gas`值
	gasLimit := uint64(21000)
	// 客户端提供SuggestGasPrice函数，用于根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	//5.转换接收方的地址信息
	toEthAddress := common.HexToAddress(toAddress)

	//5.生成未签名的以太坊事务
	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toEthAddress,
		Value:    transferAmount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     make([]byte, 0),
	}
	tx := types.NewTx(legacyTx)

	//6.使用发起人的私钥对生成的事务进行签名操作
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Network ID = (%v)", chainID))

	//这里是基于 `EIP155`,对其进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	//7.发送事务提交操作，主要就是将已签名的事物广播到整个网络中。
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}
	return nil
}
