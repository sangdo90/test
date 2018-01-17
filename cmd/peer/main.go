package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/smartm2m/blockchain/core"
)

// TODO : Jongseok
func main() {
	bc := core.NewBlockChain()
	genesisBlock := core.NewBlock()
	bc.AddBlock(genesisBlock)
	console()
}

func console() {
	for reader := bufio.NewReader(os.Stdin); ; {
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		_ = cmd
	}
}
