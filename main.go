package main

import (
	"fmt"
	"github.com/SamanNsr/cryptochain/utils/hash_utils"
)

func main() {
hash := hash_utils.CryptoHash([]byte("salamafs"), []byte("sala{{{m"), []byte("saaalam"), []byte("saaalam"))
	fmt.Println("&hash")
fmt.Println(hash)
}
