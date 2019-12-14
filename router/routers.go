package router

import (
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
