package blockchain

import (
	"fmt"
	"github.com/SamanNsr/cryptochain/utils/hash_utils"
)

type Blockchain struct {
	chain []*Block
}

func Init() *Blockchain {
	return &Blockchain{
		[]*Block{Genesis()},
	}
}

func (bc *Blockchain) AddGenesis() {
	chain := []*Block{Genesis()}
	bc.chain = chain
}

func (bc *Blockchain) AddBlock(data string) {
	newBlock := MineBlock(bc.chain[len(bc.chain)-1], data)
	bc.chain = append(bc.chain, newBlock)
}

func IsValidChain(chain []*Block) bool {
	if chain[0] != Genesis() {
		return false
	}

	statusChannel := make(chan bool, len(chain))

	for i := 0; i < len(chain); i++ {
		go func(i int) {
			block := chain[i]
			actualLastHash := chain[i-1].hash

			if block.lastHash != actualLastHash {
				statusChannel <- false
			} else {
				statusChannel <- true
			}
		}(i)

		validatedHash := hash_utils.CryptoHash(
			chain[i].timestamp, chain[i].lastHash, chain[i].data, chain[i].nonce, chain[i].difficulty,
		)

		if chain[i].hash != validatedHash {
			return false
		}

		for i := 0; i < len(chain); i++ {
			status := <-statusChannel
			if status == false {
				break
			}
		}
	}
	return true
}

func (bc *Blockchain) ReplaceChain(chain []*Block) {
	if len(chain) <= len(bc.chain) {
		fmt.Errorf("oops, incoming chain must be longer")
		return
	}
	if IsValidChain(chain) {
		fmt.Errorf("oops, incoming chain must be valid")
		return
	}

	fmt.Println("Replacing chain with")
	bc.chain = chain
}
