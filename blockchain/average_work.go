package blockchain

import (
	"fmt"
)

func Test() {
	var bc *Blockchain = Init()

	bc.AddBlock("initial")

	var times []int64
	var average int64
	var totalTime int64
	for i := 0; i < 1000; i++ {
		prevTimestamp := bc.chain[len(bc.chain)-1].timestamp

		bc.AddBlock(fmt.Sprintf("Hi, %d! Welcome.", i))
		nextBlock := bc.chain[len(bc.chain)-1]

		nextTimestamp := nextBlock.timestamp
		timeDiff := nextTimestamp.Sub(prevTimestamp).Milliseconds()
		times = append(times, timeDiff)

		totalTime = totalTime + timeDiff
		average = totalTime / int64(len(times))

		fmt.Printf(
			"The time to mine the block is: %d. Difficulty: %d. Average time: %d ms \n",
			timeDiff, nextBlock.difficulty, average,
		)
	}

}
