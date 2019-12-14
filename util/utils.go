package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type UUID []byte

var rander = rand.Reader // random function

func randomBits(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err.Error()) // rand should never fail
	}
}
func NewRandom() UUID {
	uuid := make([]byte, 16)
	randomBits([]byte(uuid))
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}

func GetDataPath() string {
	mspath := os.Getenv("MSDATA_PATH")
	if mspath == "" {
		homepath := os.Getenv("HOME")
		return fmt.Sprintf("%s/msdata/", homepath)
	}
	return mspath
}

func GetBcosNetwork() string {
	bcosurl := os.Getenv("BCOS_NETWORK")
	if bcosurl == "" {
		return "http://localhost:8545"
	}
	return bcosurl
}

func GetTokenPath() string {
	tokenpath := os.Getenv("TOKEN_PATH")
	if tokenpath == "" {
		return fmt.Sprintf("%s/msdata/tokens.json", os.Getenv("HOME"))
	}
	return tokenpath
}
