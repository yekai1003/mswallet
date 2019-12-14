//并非在GOPATH目录 -- go mod
package main

import (
	"fmt"
	"os"

	"github.com/yekai1003/mswallet/client"
)

func main() {
	url := os.Getenv("ETH_URL")
	if url == "" {
		url = "http://127.0.0.1:8545"
	}
	fmt.Println("url:", url)
	cli := client.NewCLI("./data/", url, "tokens.json")
	cli.Run()
}
