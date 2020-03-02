package main

import (
	"fmt"

	"github.com/yekai1003/mswallet/bcos"
)

const adminkey = `{"address":"3f8712acd6ed891ec329fd5ae0a93dd713237e5d","crypto":{"cipher":"aes-128-ctr","ciphertext":"623b85925792e49ac809f474c96a6dc46080d865e5fe1fa89df6c3410fbbfda1","cipherparams":{"iv":"4f0521483a5577b1573f0f63d88b0ede"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"c8ac5e6ee11526b43c2b66a44d0c0bd006fdaff23d22bd64e968406f61e38244"},"mac":"5fd86fc981d37bda5fdab0374db7916244b3dbb3eb71e92b9b6e509e21f9f009"},"id":"2785cb09-649d-4deb-88d2-de152eb78bd5","version":3}`

//0x803C0bF741ec68dA8351EE2995C0790036D244Ba
func main() {
	//bcos.DeployToken("123", "ykc")
	//bcos.TokenMint("123", "0x803C0bF741ec68dA8351EE2995C0790036D244Ba", "0x3f8712acd6ed891ec329fd5ae0a93dd713237e5d", 10000)
	//bcos.TokenBalance("0x803C0bF741ec68dA8351EE2995C0790036D244Ba", "0x3f8712acd6ed891ec329fd5ae0a93dd713237e5d")
	addr, _ := bcos.DeployAdmin("123", "yyy")
	fmt.Println(addr)
	//bcos.AdminTokenMint("123", "0xD7750a6a627A6905234821386C2beC29aeD98E04", "0x3f8712acd6ed891ec329fd5ae0a93dd713237e5d", 10000)
	// addr, _ := bcos.GetTokenAddr("0xD7750a6a627A6905234821386C2beC29aeD98E04", "0x3f8712acd6ed891ec329fd5ae0a93dd713237e5d")
	// fmt.Println(addr)
	// //0x24c8A8a1385eE3eF426AFaA02907BE8AE05DB144
	// bcos.TokenBalance(addr, "0x3f8712acd6ed891ec329fd5ae0a93dd713237e5d")
}
