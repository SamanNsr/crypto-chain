package blockchain

type Blockchain struct {
	chain []*Block
}

func (bc *Blockchain) AddGenesis() {
	bc.chain = append(bc.chain, Genesis())
}

func (bc *Blockchain) AddBlock(data string) {
	var newBlock *Block
	newBlock.MineBlock(bc.chain[len(bc.chain)-1], data)

	bc.chain = append(bc.chain, newBlock)
}

//concurrency may needed
func IsValidChain(chain []*Block) bool {
	if chain[0] != Genesis() {
		return false
	}

	for i := 0; i < len(chain); i++ {
		block := chain[i]
		actualLastHash := chain[i-1].hash

		if block.lastHash != actualLastHash {
			return false
		}

	}

	return true
}
