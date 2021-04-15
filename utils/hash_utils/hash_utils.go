package hash_utils

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func GetBytes(key interface{}) ([]byte, error) {
	buf, ok := key.([]byte)
	if !ok {
		fmt.Println("err")
		return nil, fmt.Errorf("ooops, did not work")
	}

	return buf, nil
}

func CryptoHash(objs ...interface{}) string {
	digester := crypto.SHA256.New()
	for _, ob := range objs {
		fmt.Fprint(digester, ob)
	}
	sha := base64.URLEncoding.EncodeToString(digester.Sum(nil))
	return sha
}