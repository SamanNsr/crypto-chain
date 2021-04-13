package block

import "fmt"

type Block struct {
	timestamp string
	lastHash string
	hash string
	data string
}

func _init_(timestamp string, lastHash string , hash string, data string) *Block {
	return &Block{timestamp, lastHash, hash, data}
}

func main() {
	var block1 *Block
	block1 = _init_("01/01/01", "foo-lastHash", "foo-hash", "foo-data")
	fmt.Print(*block1)
}