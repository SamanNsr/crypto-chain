package blockchain

import "github.com/SamanNsr/cryptochain/block"

type Blockchain struct {
	chain []*block.Block
}

func (bc *Blockchain) AddGenesis() {
	bc.chain = append(bc.chain, block.Genesis())
}