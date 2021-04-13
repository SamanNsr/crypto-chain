package block

import (
	"github.com/SamanNsr/cryptochain/utils/date_utils"
	"github.com/SamanNsr/cryptochain/utils/hash_utils"
)

type Block struct {
	timestamp string
	lastHash []byte
	hash []byte
	data interface{}
}

func Genesis() *Block{
	return &Block{
		 "01/01/01",
		   hash_utils.CryptoHash("foo-lastHash"),
		     hash_utils.CryptoHash("foo-hash"),
		     "foo-data",
		}
}

func (b *Block) NewBlock(timestamp string, lastHash []byte , hash []byte, data string) {
	b.timestamp = timestamp
	b.lastHash = lastHash
	b.hash = hash
	b.data = data
}

func (b *Block) MineBlock(lastBlock *Block, data string) {
	timestamp := date_utils.GetNowString()
	lastHash := lastBlock.hash
 	b.timestamp = timestamp
 	b.lastHash = lastHash
 	b.data = data
 	b.hash = hash_utils.CryptoHash(timestamp ,lastHash, data)
}