package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/yekai1003/mswallet/util"

	"github.com/yekai1003/mswallet/client"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type MintMsg struct {
	PassWord string `json:"password"`
	ToAddr   string `json:"toaddr"`
	Val      uint   `json:"value"`
}

type BurnMsg struct {
	PassWord string `json:"password"`
	Owner    string `json:"owner"`
	Val      uint   `json:"value"`
}

func ResponseData(c echo.Context, resp *util.Resp) {
	resp.ErrMsg = util.RecodeText(resp.Errno)
	c.JSON(http.StatusOK, resp)
}

// Handler
func Hello(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

//注册功能
//post: /register {username,password}
func Register(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp util.Resp
	resp.Errno = util.RECODE_OK
	defer ResponseData(c, &resp)

	//2. 解析数据
	user := User{}
	if err := c.Bind(&user); err != nil {
		fmt.Println("Register:Failed to Bind ", err, user)
		resp.Errno = util.RECODE_REGERR
		return err
	}
	//3. 根据请求的数据创建账户
	cli := client.NewCLI(util.GetDataPath(), util.GetBcosNetwork(), util.GetTokenPath())
	acc, err := cli.CreateWallet(user.UserName, user.PassWord)
	if err != nil {
		fmt.Println("Failed to CreateWallet", err)
		resp.Errno = util.RECODE_CREATEWALLETERR
		return err
	}
	resp.Data = acc
	return nil
}

//登陆功能
//post: /login {username,password}
func Login(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp util.Resp
	resp.Errno = util.RECODE_OK
	defer ResponseData(c, &resp)

	//2. 解析数据
	user := User{}
	if err := c.Bind(&user); err != nil {
		fmt.Println("Login:Failed to Bind ", err, user)
		resp.Errno = util.RECODE_REGERR
		return err
	}
	//3. 根据请求的数据创建账户
	cli := client.NewCLI(util.GetDataPath(), util.GetBcosNetwork(), util.GetTokenPath())
	addr, err := cli.Login(user.UserName, user.PassWord)
	if err != nil {
		fmt.Println("Loin err", err)
		resp.Errno = util.RECODE_LOGINERR
	}
	resp.Data = addr
	return nil
}

//post: /mint {toname,password,value}
func TokenMint(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp util.Resp
	resp.Errno = util.RECODE_OK
	defer ResponseData(c, &resp)

	//2. 解析数据
	mintMsg := MintMsg{}
	if err := c.Bind(&mintMsg); err != nil {
		fmt.Println("TokenMint:Failed to Bind ", err, mintMsg)
		resp.Errno = util.RECODE_ACCERR
		return err
	}

	//3. 根据请求的数据创建账户
	cli := client.NewCLI(util.GetDataPath(), util.GetBcosNetwork(), util.GetTokenPath())
	hash, err := cli.AdminMintToken(mintMsg.PassWord, mintMsg.ToAddr, int64(mintMsg.Val))
	if err != nil {
		fmt.Println("TokenMint:Failed to AdminMintToken ", err, mintMsg)
		resp.Errno = util.RECODE_BCOSERR
		return err
	}
	resp.Data = hash
	return err
}

//post: /tokenburn {password,value}
func TokenBurn(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp util.Resp
	resp.Errno = util.RECODE_OK
	defer ResponseData(c, &resp)

	//2. 解析数据
	burnMsg := BurnMsg{}
	if err := c.Bind(&burnMsg); err != nil {
		fmt.Println("TokenBurn:Failed to Bind ", err, burnMsg)
		resp.Errno = util.RECODE_UNKNOWERR
		return err
	}

	//3. 根据请求的数据创建账户
	cli := client.NewCLI(util.GetDataPath(), util.GetBcosNetwork(), util.GetTokenPath())
	hash, err := cli.AdminBurnToken(burnMsg.PassWord, burnMsg.Owner, int64(burnMsg.Val))
	if err != nil {
		fmt.Println("TokenBurn:Failed to AdminBurnToken ", err, burnMsg)
		resp.Errno = util.RECODE_BCOSERR
		return err
	}
	resp.Data = hash
	return err
}

//GET: /balance/:fromaddr
func GetBalance(c echo.Context) error {
	//1. 响应数据结构初始化
	fmt.Println("begin GetBalance")
	var resp util.Resp
	resp.Errno = util.RECODE_OK
	defer ResponseData(c, &resp)

	//2. 解析数据
	fromaddr := c.Param("fromaddr")

	if fromaddr == "" {
		fmt.Println("TokenBurn:Failed to get  fromaddr", fromaddr)
		resp.Errno = util.RECODE_UNKNOWERR
		return errors.New("no fromaddr")
	}
	fmt.Println(fromaddr)
	//3. 根据请求的数据创建账户
	cli := client.NewCLI(util.GetDataPath(), util.GetBcosNetwork(), util.GetTokenPath())
	val, err := cli.ADTokenBalance(fromaddr)
	if err != nil {
		fmt.Println("TokenBurn:Failed to ADTokenBalance ", err, fromaddr)
		resp.Errno = util.RECODE_BCOSERR
		return err
	}
	resp.Data = struct {
		Balance int64 `json:"balance"`
	}{Balance: val.Int64()}
	fmt.Println(resp)
	return err
}
