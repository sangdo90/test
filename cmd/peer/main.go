package main

import (
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/log"
)

// TODO : Jongseok
func main() {
	bc := core.NewBlockChain()
	genesisBlock := core.NewBlock()
	bc.AddBlock(genesisBlock)

	log.SetLogLevel(log.DebugLogLevel)
	console.Start()
}

// func console() {
// 	for reader := bufio.NewReader(os.Stdin); ; {
// 		fmt.Print("> ")
// 		cmd, _ := reader.ReadString('\n')
// 		_ = cmd
// 	}
// }
