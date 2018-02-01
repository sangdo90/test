package main

import (
	"github.com/smartm2m/blockchain/execute"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// TODO : Jongseok
func main() {
	//bc := core.AppendBlockchain()
	// TODO: Using a blockchain
	//_ = bc

	log.SetLogLevel(log.DebugLogLevel)
	//console.RegisterBlockchain(bc, core.AppendBlockchain)
	//console.RegisterBlock(core.NewBlock)
	execute.BlockchainCommands()
	execute.ConsensusCommands()
	RegisterCommand()
	console.Start()
}

// RegisterCommand register commands for manage blockchains.
func RegisterCommand() {
	_ = command.AddCommand("", command.Command{
		Name:        "quit",
		ShortName:   "q",
		Description: "Exit the program",
		Commands:    nil,
		Flags:       nil,
		Run: func() error {
			console.GetContext().Quit()
			return nil
		},
	})
}
