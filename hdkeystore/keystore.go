package hdkeystore

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yekai1003/mswallet/util"

	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/accounts/keystore"
	"github.com/yekai1003/gobcos/common"
	"github.com/yekai1003/gobcos/crypto"
)

type HDkeyStore struct {
	keysDirPath string
	scryptN     int
	scryptP     int
	Key         keystore.Key
}

//给出一个生成HDkeyStore对象的方法,通过privatekey生成
func NewHDkeyStore(path string, privateKey *ecdsa.PrivateKey) *HDkeyStore {
	uuid := []byte(util.NewRandom())
	if privateKey == nil {
		return &HDkeyStore{
			keysDirPath: path,
			scryptN:     keystore.LightScryptN,
			scryptP:     keystore.LightScryptP,
			Key:         keystore.Key{},
		}
	}
	key := keystore.Key{
		Id:         uuid,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}
	return &HDkeyStore{
		keysDirPath: path,
		scryptN:     keystore.LightScryptN,
		scryptP:     keystore.LightScryptP,
		Key:         key,
	}
}

func (ks *HDkeyStore) GetKey(addr common.Address, filename, auth string) (*keystore.Key, error) {
	// Load the key from the keystore and decrypt its contents

	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, addr)
	}
	ks.Key = *key
	return key, nil
}

func (ks HDkeyStore) StoreKey(filename string, key *keystore.Key, auth string) error {
	keyjson, err := keystore.EncryptKey(key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	return WriteKeyFile(filename, keyjson)
}

func (ks HDkeyStore) JoinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}

	return filepath.Join(ks.keysDirPath, filename)
}

func WriteKeyFile(file string, content []byte) error {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()

	return os.Rename(f.Name(), file)
}

func (ks HDkeyStore) NewTransactOpts() *bind.TransactOpts {
	return bind.NewKeyedTransactor(ks.Key.PrivateKey)
}
