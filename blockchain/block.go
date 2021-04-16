package blockchain

import (
	"github.com/SamanNsr/cryptochain/utils/date_utils"
	"github.com/SamanNsr/cryptochain/utils/hash_utils"
	"strings"
	"time"
)

type Block struct {
	timestamp  time.Time
	lastHash   string
	hash       string
	data       string
	nonce      int
	difficulty int
}

const InitialDifficulty = 3
const mineRate = 1000

func Genesis() *Block {
	return &Block{
		date_utils.GetNow(),
		hash_utils.CryptoHash("foo-lastHash"),
		hash_utils.CryptoHash("foo-hash"),
		"foo-data",
		0,
		InitialDifficulty,
	}
}

func NewBlock(timestamp time.Time, lastHash string, hash string, data string, nonce int, difficulty int) *Block {
	return &Block{
		timestamp,
		lastHash,
		hash,
		data,
		nonce,
		difficulty,
	}
}

func MineBlock(lastBlock *Block, data string) *Block {
	var hash string
	var timestamp time.Time
	lastHash := lastBlock.hash
	difficulty := lastBlock.difficulty
	nonce := 0

	for founded := false; !founded; founded = hash[0:difficulty] == strings.Repeat("0", difficulty) {
		nonce++
		timestamp = date_utils.GetNow()
		difficulty = adjustDifficulty(lastBlock, date_utils.GetNow())
		hash = hash_utils.HexToBinary(
			hash_utils.CryptoHash(timestamp, lastHash, data, nonce, difficulty),
		)
	}
	return NewBlock(timestamp, lastHash, hash, data, nonce, difficulty)
}

func adjustDifficulty(originalBlock *Block, timestamp time.Time) int {
	difficulty := originalBlock.difficulty

	if difficulty < 1 {
		return 1
	}
	difference := timestamp.Sub(originalBlock.timestamp).Milliseconds()

	if difference > mineRate {
		return difficulty - 1
	}

	return difficulty + 1
}
