package main

import (
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/console/command"
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
	_ = command.AddCommand("", command.Command{
		Name:        "quit",
		Description: "Exit the program",
		Commands:    nil,
		Flags:       nil,
		Run: func(args []string) error {
			console.GetContext().Quit()
			return nil
		},
	})
}
