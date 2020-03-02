package bcos

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/client"
	"github.com/yekai1003/gobcos/common"
)

const adminkey = `{"address":"3f8712acd6ed891ec329fd5ae0a93dd713237e5d","crypto":{"cipher":"aes-128-ctr","ciphertext":"623b85925792e49ac809f474c96a6dc46080d865e5fe1fa89df6c3410fbbfda1","cipherparams":{"iv":"4f0521483a5577b1573f0f63d88b0ede"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"c8ac5e6ee11526b43c2b66a44d0c0bd006fdaff23d22bd64e968406f61e38244"},"mac":"5fd86fc981d37bda5fdab0374db7916244b3dbb3eb71e92b9b6e509e21f9f009"},"id":"2785cb09-649d-4deb-88d2-de152eb78bd5","version":3}`

func DeployToken(pass, tokenName string) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 合约部署
	//common.Address, *types.RawTransaction, *Cert, error
	address, tx, _, err := DeployErc20(auth, cli, tokenName)
	if err != nil {
		log.Panic("Failed to DeployCert", err)
	}
	fmt.Println(tx.Hash().Hex(), address.Hex())
}
func TokenMint(pass, contract, toAddr string, val int64) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	token, err := NewErc20(common.HexToAddress(contract), cli)
	if err != nil {
		log.Panic("Failed to NewErc20", err)
	}
	tx, err := token.Mint(auth, common.HexToAddress(toAddr), big.NewInt(val))
	if err != nil {
		log.Panic("Failed to DeployCert", err)
	}
	fmt.Println(tx.Hash().Hex())
}

// func BurnToken(pass, contract string, val int64) {
// 	//1. 连接到fisco节点
// 	cli, err := client.Dial("http://10.211.55.3:8545", 1)
// 	if err != nil {
// 		log.Panic("Failed to Dial", err)
// 	}
// 	//延迟关闭连接
// 	defer cli.Close()

// 	//2.签名
// 	keyin := strings.NewReader(adminkey)
// 	auth, err := bind.NewTransactor(keyin, pass)
// 	//3. 调用函数
// 	token, err := NewErc20(common.HexToAddress(contract), cli)
// 	if err != nil {
// 		log.Panic("Failed to NewErc20", err)
// 	}
// 	tx, err := token.Burn(auth, big.NewInt(val))
// 	if err != nil {
// 		log.Panic("Failed to Burn", err)
// 	}
// 	fmt.Println(tx.Hash().Hex())
// }

func TokenBalance(contract, fromaddr string) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2. 构造合约实例
	token, err := NewErc20(common.HexToAddress(contract), cli)
	if err != nil {
		log.Panic("Failed to DeployCert", err)
	}
	//3. 验证用户证书
	opts := bind.CallOpts{
		From: common.HexToAddress(fromaddr),
	}

	//调用验证
	data, err := token.BalanceOf(&opts, opts.From)
	if err != nil {
		log.Panic("Failed to GetUserCerts", err)
	}
	fmt.Println(data)
}

func DeployAdmin(pass, tokenName string) (string, error) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3.部署管理合约
	address, tx, _, err := DeployTokenadmin(auth, cli, tokenName)
	if err != nil {
		log.Panic("Failed to DeployTokenadmin")
	}
	fmt.Println(tx.Hash().Hex(), address.Hex())
	return address.Hex(), err
}

func GetTokenAddr(contract, fromaddr string) (string, error) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()
	//2. 构造实例
	ta, err := NewTokenadmin(common.HexToAddress(contract), cli)
	if err != nil {
		fmt.Println("Failed to NewTokenadmin", err)
		return "", err
	}
	//3. 调用获得地址
	opts := bind.CallOpts{
		From: common.HexToAddress(fromaddr),
	}
	address, err := ta.Tokenaddr(&opts)
	if err != nil {
		fmt.Println("Failed to Tokenaddr", err)
		return "", err
	}
	return address.Hex(), err

}

func AdminTokenMint(pass, contract, toAddr string, val int64) error {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		fmt.Println("Failed to Dial", err)
		return err
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	ta, err := NewTokenadmin(common.HexToAddress(contract), cli)
	if err != nil {
		fmt.Println("Failed to NewTokenadmin", err)
		return err
	}
	tx, err := ta.Mint(auth, common.HexToAddress(toAddr), big.NewInt(val))
	if err != nil {
		log.Panic("Failed to Mint", err)
	}
	fmt.Println(tx.Hash().Hex())
	return err
}

func AdminTokenBurn(pass, contract, owner string, val int64) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	ta, err := NewTokenadmin(common.HexToAddress(contract), cli)
	if err != nil {
		log.Panic("Failed to NewTokenadmin", err)
	}
	tx, err := ta.Burn(auth, common.HexToAddress(owner), big.NewInt(val))
	if err != nil {
		log.Panic("Failed to Burn", err)
	}
	fmt.Println(tx.Hash().Hex())
}

func AdminTokenTransfer(pass, contract, toAddr string, val int64) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	ta, err := NewTokenadmin(common.HexToAddress(contract), cli)
	if err != nil {
		log.Panic("Failed to NewTokenadmin", err)
	}
	tx, err := ta.Transfer(auth, common.HexToAddress(toAddr), big.NewInt(val))
	if err != nil {
		log.Panic("Failed to Transfer", err)
	}
	fmt.Println(tx.Hash().Hex())
}

func AdminTokenUpgrade(pass, contract, tokenAddr string) {
	//1. 连接到fisco节点
	cli, err := client.Dial("http://10.211.55.3:8545", 1)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	ta, err := NewTokenadmin(common.HexToAddress(contract), cli)
	if err != nil {
		log.Panic("Failed to NewTokenadmin", err)
	}
	tx, err := ta.Upgrade(auth, common.HexToAddress(tokenAddr))
	if err != nil {
		log.Panic("Failed to Upgrade", err)
	}
	fmt.Println(tx.Hash().Hex())
}
