package main

import (
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/blockchain/execute"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/log"
)

// TODO : Jongseok
func main() {
	bc := core.NewBlockChain()

	// TODO: Using a blockchain
	_ = bc

	log.SetLogLevel(log.DebugLogLevel)
	console.RegisterBlockchain(bc, core.NewBlockChain)
	console.RegisterBlock(core.NewBlock)
	RegisterCommand()
	console.Start()
}

// RegisterCommand register commands for manage blockchains.
func RegisterCommand() {
	execute.BlockchainCommands()
}

// func console() {
// 	for reader := bufio.NewReader(os.Stdin); ; {
// 		fmt.Print("> ")
// 		cmd, _ := reader.ReadString('\n')
// 		_ = cmd
// 	}
// }
