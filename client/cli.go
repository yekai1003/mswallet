package client

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/yekai1003/mswallet/bcos"
	"github.com/yekai1003/mswallet/hd"
	"github.com/yekai1003/mswallet/hdkeystore"

	"log"
	"os"

	"github.com/howeyc/gopass"
	"github.com/tyler-smith/go-bip39"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/client"
	"github.com/yekai1003/gobcos/common"
)

const ADMIN_CONTRACT = "ADMIN_CONTRACT"
const ADMIN_JSON = "admin.json"
const ADMIN_CONTRACT_PRE = "0x5Fd877666843cbB8D4349C29069033aA3C541a50"

type CLI struct {
	DataPath          string
	NetWorkUrl        string
	TokenFile         string
	AdminContractAddr string
}

type TokenConfig struct {
	Symbol string `json:"symbol"`
	Addr   string `json:"addr"`
}

var AdminKey = `{"address":"3f8712acd6ed891ec329fd5ae0a93dd713237e5d","crypto":{"cipher":"aes-128-ctr","ciphertext":"623b85925792e49ac809f474c96a6dc46080d865e5fe1fa89df6c3410fbbfda1","cipherparams":{"iv":"4f0521483a5577b1573f0f63d88b0ede"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"c8ac5e6ee11526b43c2b66a44d0c0bd006fdaff23d22bd64e968406f61e38244"},"mac":"5fd86fc981d37bda5fdab0374db7916244b3dbb3eb71e92b9b6e509e21f9f009"},"id":"2785cb09-649d-4deb-88d2-de152eb78bd5","version":3}`

func init() {
	fd, err := os.Open(ADMIN_JSON)
	if err != nil {
		fmt.Println("Failed to open file ", err, ADMIN_JSON)
		//AdminKeyIn = strings.NewReader(AdminKey)
		return
	}
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("Failed to ReadAll ", err, ADMIN_JSON)
		//AdminKeyIn = strings.NewReader(AdminKey)
		return
	}
	AdminKey = string(data)
}

func NewCLI(path, url, tokenfile string) *CLI {
	adAddr := os.Getenv(ADMIN_CONTRACT)
	if adAddr == "" {
		adAddr = ADMIN_CONTRACT_PRE
	}
	return &CLI{
		DataPath:          path,
		NetWorkUrl:        url,
		TokenFile:         tokenfile,
		AdminContractAddr: adAddr,
	}
}

//提供帮助
func (c CLI) Help() {
	fmt.Println("./bcoswallet createwallet -name ACCOUNT_NAME --for create new wallet")
	fmt.Println("./bcoswallet balance -name ACCOUNT_NAME  --for get balance")
	fmt.Println("./bcoswallet sendtoken -name ACCOUNT_NAME -symbol SYMBOL -toaddr ADDRESS -amount AMOUNT --for get balance")
	fmt.Println("./bcoswallet addtoken -addr CONTRACT_ADDR --for add token to wallet")

}

//参数检测
func (c CLI) Valid() {
	if len(os.Args) < 2 {
		c.Help()
		os.Exit(-1)
	}
}

func (c CLI) Run() {
	//运行前先检测
	c.Valid()
	//(一) 创建账户
	//解析命令行
	//1. 分类设定
	createwalletCMD := flag.NewFlagSet("createwallet", flag.ExitOnError)
	//2. 设定要解析具体参数
	createwalletCMD_name := createwalletCMD.String("name", "yekai", "ACCOUNT_NAME")

	//(三) 查询token余额
	balanceCMD := flag.NewFlagSet("balance", flag.ExitOnError)
	balanceCMD_name := balanceCMD.String("name", "yekai", "ACCOUNT_NAME")
	//(四) 添加token
	addtokenCMD := flag.NewFlagSet("addtoken", flag.ExitOnError)
	addtokenCMD_addr := addtokenCMD.String("addr", "", "CONTRACT_ADDR")
	//(五) token转账
	sendtokenCMD := flag.NewFlagSet("sendtoken", flag.ExitOnError)
	sendtokenCMD_name := sendtokenCMD.String("name", "", "ACCOUNT_NAME")
	sendtokenCMD_symbol := sendtokenCMD.String("symbol", "", "SYMBOL")
	sendtokenCMD_toaddr := sendtokenCMD.String("toaddr", "", "ADDRESS")
	sendtokenCMD_amount := sendtokenCMD.Int64("amount", 0, "AMOUNT")

	switch os.Args[1] {
	case "createwallet":
		//真的解析参数
		err := createwalletCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Failed to createwalletCMD.Parse", err)
			return
		}
	case "balance":
		err := balanceCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Failed to balanceCMD.Parse", err)
			return
		}
	case "addtoken":
		err := addtokenCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Failed to addtokenCMD.Parse", err)
			return
		}
	case "sendtoken":
		err := sendtokenCMD.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Failed to sendtokenCMD.Parse", err)
			return
		}
	default:
		c.Help()
		os.Exit(1)
	}

	//运行具体功能
	if createwalletCMD.Parsed() {
		//创建账户
		fmt.Println(*createwalletCMD_name)
		if *createwalletCMD_name == "" {
			fmt.Println("createwalletCMD_name can not null")
			return
		}
		//解决密码问题
		pass, _ := gopass.GetPasswd()
		c.CreateWallet(*createwalletCMD_name, string(pass))
	}

	//查询余额
	if balanceCMD.Parsed() {
		if *balanceCMD_name == "" {
			fmt.Println("balanceCMD_name can not null")
			return
		}
		//查询余额
		c.GetTokensBalance(*balanceCMD_name)
	}
	//添加token
	if addtokenCMD.Parsed() {
		if *addtokenCMD_addr == "" {
			fmt.Println("contract addr can not null")
			return
		}
		//代码实现
		c.AddToken(*addtokenCMD_addr)
	}
	// 发送token
	if sendtokenCMD.Parsed() {
		if *sendtokenCMD_symbol == "" || *sendtokenCMD_toaddr == "" || *sendtokenCMD_amount <= 0 {
			fmt.Println("sendtokenCMD params error")
			return
		}
		//代码实现
		c.SendToken(*sendtokenCMD_name, *sendtokenCMD_toaddr, *sendtokenCMD_symbol, *sendtokenCMD_amount)
	}
}

//创建账户功能 name = yekai  /data/yekai/0x------- /data/xiaohong/0x
func (c CLI) CreateWallet(name, pass string) (string, error) {
	//目录检测
	if c.getAddr(name) != "" {
		fmt.Printf("Acct %s is exists\n", name)
		return "", errors.New("Acct" + name + "exists")
	}

	//1. NewEntropy 必须是32的整数倍，并且在128- 256之间

	entropy, _ := bip39.NewEntropy(128)
	//2. 助记词
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(mnemonic)

	wallet, err := hd.NewFromMnemonic(mnemonic, "")
	if err != nil {
		fmt.Println("Failed to NewFromMnemonic", err)
		return "", err
	}
	path, _ := hd.ParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		fmt.Println("Failed to Derive", err)
		return "", err
	}
	fmt.Println(account.Address.Hex())
	//PrivateKey(account accounts.Account) (*ecdsa.PrivateKey, error)
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		fmt.Println("Failed to PrivateKey", err)
		return "", err
	}
	//调用生成keystore对象

	hdks := hdkeystore.NewHDkeyStore(c.DataPath+name, privateKey)
	//StoreKey(filename string, key *Key, auth string)
	err = hdks.StoreKey(hdks.JoinPath(account.Address.Hex()), &hdks.Key, pass)
	if err != nil {
		fmt.Println("Failed to StoreKey", err)
		return "", err
	}
	return account.Address.Hex(), nil
}

func (c CLI) getAddr(name string) string {
	infos, err := ioutil.ReadDir(c.DataPath + name)
	if err != nil {
		fmt.Println("Failed to ReadDir ", err)
		return ""
	}
	for _, v := range infos {
		if !v.IsDir() {
			if strings.HasPrefix(v.Name(), "0x") {
				return v.Name()
			}
		}
	}
	return ""
}

func hex2bigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}

//添加token

func (c CLI) AddToken(addr string) {
	//0. 读取配置文件
	tokens := c.ReadToken()

	//0.1 校验没有重复的地址
	if c.CheckToken(addr, tokens) {
		fmt.Println("token is exists", addr)
		return
	}
	//1. 连接到网络
	client, err := client.Dial(c.NetWorkUrl, 1)
	if err != nil {
		log.Panic("Failed to client.Dial", err)
	}
	//2. 通过合约地址创建合约实例
	ins, err := bcos.NewErc20(common.HexToAddress(addr), client)
	if err != nil {
		log.Panic("Failed to erc20.NewErc20:", err)
	}
	opts := bind.CallOpts{
		From: common.HexToAddress(addr),
	}
	sym, err := ins.Symbol(&opts)
	if err != nil {
		log.Panic("Failed to ins.Symbol:", err)
	}
	//3. 写入配置文件
	tokens = append(tokens, TokenConfig{sym, addr})
	content, err := json.Marshal(tokens)
	if err != nil {
		log.Panic("Failed to json.Marshal:", err)
	}
	hdkeystore.WriteKeyFile(c.TokenFile, content)
}

func (c CLI) ReadToken() []TokenConfig {
	tokens := []TokenConfig{}
	data, err := ioutil.ReadFile(c.TokenFile)
	if err != nil {
		//log.Panic("Failed to ReadFile", err)
		fmt.Println("Failed to ReadFile")
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &tokens)
		if err != nil {
			log.Panic("Failed to Unmarshal", err)
		}
	}
	return tokens
}

func (c CLI) CheckToken(addr string, tokens []TokenConfig) bool {
	for _, token := range tokens {
		if addr == token.Addr {
			return true
		}
	}
	return false
}

func (c CLI) Login(acctname, password string) (string, error) {
	hdks := hdkeystore.NewHDkeyStore(c.DataPath+acctname, nil)
	fromAddr := c.getAddr(acctname)
	if fromAddr == "" {
		fmt.Println("acct not exists ")
		return "", errors.New("acct not exists")
	}
	_, err := hdks.GetKey(common.HexToAddress(fromAddr), hdks.JoinPath(fromAddr), password)
	if err != nil {
		fmt.Println("Failed to hdks.GetKey ", err)
		return "", err
	}
	return fromAddr, nil
}

func (c CLI) SendToken(acctname, toaddr, symbol string, amount int64) {
	//1. 连接到网络
	client, err := client.Dial(c.NetWorkUrl, 1)
	if err != nil {
		log.Panic("Failed to client.Dial", err)
	}
	defer client.Close()
	//2. 生成合约实例
	contract_addr := c.GetContractAddr(symbol)
	if contract_addr == "" {
		fmt.Println("symbol not exists", symbol)
		return
	}
	ins, err := bcos.NewErc20(common.HexToAddress(contract_addr), client)
	if err != nil {
		log.Panic("Failed to erc20.NewErc20", err)
	}
	//3. 设置签名
	hdks := hdkeystore.NewHDkeyStore(c.DataPath+acctname, nil)
	fromAddr := c.getAddr(acctname)
	_, err = hdks.GetKey(common.HexToAddress(fromAddr), hdks.JoinPath(fromAddr), "123")
	if err != nil {
		log.Panic("Failed to hdks.GetKey ", err)
	}
	opts := hdks.NewTransactOpts()
	//4. 合约调用
	//opts *bind.TransactOpts, to common.Address, value *big.Int
	value := big.NewInt(amount)
	_, err = ins.Transfer(opts, common.HexToAddress(toaddr), value)
	if err != nil {
		log.Panic("Failed to ins.Transfer ", err)
	}
}

func (c CLI) GetContractAddr(symbol string) string {
	data, err := ioutil.ReadFile(c.TokenFile)
	if err != nil {
		log.Panic("Failed ioutil.ReadFile ", err)
	}
	tokens := []TokenConfig{}
	err = json.Unmarshal(data, &tokens)
	if err != nil {
		log.Panic("Failed json.Unmarshal ", err)
	}

	for _, v := range tokens {
		if v.Symbol == symbol {
			return v.Addr
		}
	}

	return ""

}

func (cli *CLI) GetTokensBalance(acct_name string) {
	//先读取要获取哪些token
	data, err := ioutil.ReadFile(cli.TokenFile)
	if err != nil {
		log.Panic("failed to GetTokensBalance when ReadFile ", err)
	}
	tokens := []TokenConfig{}
	err = json.Unmarshal(data, &tokens)
	if err != nil {
		log.Panic("failed to GetTokensBalance when Unmarshal ", err)
	}
	//再读取账户下有哪些地址
	acctaddr := cli.getAddr(acct_name)

	opts := bind.CallOpts{
		From: common.HexToAddress(acctaddr),
	}

	//每种token，每个地址做一个处理
	for _, token := range tokens {
		amount, err := cli.getTokenBalance(token.Addr, acctaddr, &opts)
		if err != nil {
			fmt.Println("failed to getTokenBalance: ", err)
			continue
		}
		fmt.Printf("%s:%v\n", token.Symbol, amount)
	}
}

func (cli *CLI) getTokenBalance(contract, account string, opts *bind.CallOpts) (*big.Int, error) {
	client, err := client.Dial(cli.NetWorkUrl, 1)
	if err != nil {
		fmt.Println("failed to getTokenBalance when dial", err)
		return nil, err
	}
	ins, err := bcos.NewErc20(common.HexToAddress(contract), client)
	if err != nil {
		fmt.Println("failed to getTokenBalance when NewErc20", err)
		return nil, err
	}
	return ins.BalanceOf(opts, common.HexToAddress(account))
}

//管理员挖矿给工人
func (c CLI) AdminMintToken(pass, toaddr string, amount int64) (string, error) {
	//1. 连接到网络
	client, err := client.Dial(c.NetWorkUrl, 1)
	if err != nil {
		fmt.Println("Failed to client.Dial", err)
		return "", err
	}
	defer client.Close()
	//2. 生成合约实例

	ins, err := bcos.NewTokenadmin(common.HexToAddress(c.AdminContractAddr), client)
	if err != nil {
		fmt.Println("Failed to erc20.NewTokenadmin", err)
		return "", err
	}
	//3. 设置签名
	keyin := strings.NewReader(AdminKey)
	fmt.Println(pass, AdminKey)
	auth, err := bind.NewTransactor(keyin, pass)
	if err != nil {
		fmt.Println("AdminMintToken:Failed to NewTransactor", err)
		return "", err
	}
	//4. 合约调用
	//opts *bind.TransactOpts, to common.Address, value *big.Int
	value := big.NewInt(amount)
	tx, err := ins.Mint(auth, common.HexToAddress(toaddr), value)
	if err != nil {
		fmt.Println("Failed to ins.Mint ", err)
		return "", err
	}
	return tx.Hash().Hex(), err
}

//管理员销毁token
func (c CLI) AdminBurnToken(pass, owner string, amount int64) (string, error) {
	//1. 连接到网络
	client, err := client.Dial(c.NetWorkUrl, 1)
	if err != nil {
		fmt.Println("Failed to client.Dial", err)
		return "", err
	}
	defer client.Close()
	//2. 生成合约实例

	ins, err := bcos.NewTokenadmin(common.HexToAddress(c.AdminContractAddr), client)
	if err != nil {
		fmt.Println("AdminBurnToken:Failed to bcos.NewTokenadmin", err)
		return "", err
	}
	//3. 设置签名
	//2.签名
	keyin := strings.NewReader(AdminKey)
	fmt.Println(pass, AdminKey)
	auth, err := bind.NewTransactor(keyin, pass)
	if err != nil {
		fmt.Println("Failed to NewTransactor", err)
		return "", err
	}
	//4. 合约调用
	value := big.NewInt(amount)
	tx, err := ins.Burn(auth, common.HexToAddress(owner), value)
	if err != nil {
		fmt.Println("Failed to ins.Burn ", err)
		return "", err
	}
	return tx.Hash().Hex(), err
}

//查询token余额
func (c CLI) ADTokenBalance(fromaddr string) (big.Int, error) {
	value := big.NewInt(0)
	//1. 连接到网络
	client, err := client.Dial(c.NetWorkUrl, 1)
	if err != nil {
		fmt.Println("Failed to client.Dial", err)
		return *value, err
	}
	defer client.Close()
	//2. 生成合约实例

	ins, err := bcos.NewTokenadmin(common.HexToAddress(c.AdminContractAddr), client)
	if err != nil {
		fmt.Println("Failed to bcos.NewTokenadmin", err)
		return *value, err
	}
	//3. 设置签名
	opts := bind.CallOpts{
		From: common.HexToAddress(fromaddr),
	}
	//4. 合约调用

	tokenaddr, err := ins.Tokenaddr(&opts)
	if err != nil {
		fmt.Println("Failed to ins.Tokenaddr ", err)
		return *value, err
	}
	token, err := bcos.NewErc20(tokenaddr, client)
	if err != nil {
		fmt.Println("Failed to bcos.NewErc20 ", err)
		return *value, err
	}
	value, err = token.BalanceOf(&opts, opts.From)
	if err != nil {
		fmt.Println("Failed to token.BalanceOf ", err)
		return *value, err
	}
	fmt.Println(value)
	return *value, err
}
