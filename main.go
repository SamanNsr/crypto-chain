package main

import (
	"fmt"
	"github.com/SamanNsr/cryptochain/blockchain"
	"github.com/SamanNsr/cryptochain/utils/hash_utils"
)

func main() {
hash := hash_utils.CryptoHash(("salamafs"), ("sala{{{m"), ("saaalam"), ("saaalam"), blockchain.Genesis())
	fmt.Println("&hash")
fmt.Println(hash)
}
