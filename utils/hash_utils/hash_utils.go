package hash_utils

import (
	"crypto"
	"fmt"
	"reflect"
)

func CryptoHash(objs ...interface{}) []byte {
	digester := crypto.SHA256.New()
	for _, ob := range objs {
		fmt.Fprint(digester, reflect.TypeOf(ob))
		fmt.Fprint(digester, ob)
	}
	return digester.Sum(nil)
}