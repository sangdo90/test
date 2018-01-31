package main

import (
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// TODO : Jongseok
func main() {
	bc := core.NewBlockchain()

	// TODO: Using a blockchain
	_ = bc

	log.SetLogLevel(log.DebugLogLevel)
<<<<<<< HEAD
	//console.RegisterBlockchain(bc, core.NewBlockchain)
	//console.RegisterBlock(core.NewBlock)
	execute.BlockchainCommands()
=======
	console.RegisterBlockchain(bc, core.NewBlockchain)
	console.RegisterBlock(core.NewBlock)
>>>>>>> 3498b7b3eca80b39fc2ab47b0a8596d40667c70e
	RegisterCommand()
	console.Start()
}

// RegisterCommand register commands for manage blockchains.
func RegisterCommand() {
<<<<<<< HEAD
=======
	command.AddCommand("", command.Command{
		Name:        "block",
		Description: "manage a single block.",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			log.Info("block command")
			return nil
		},
	})
	command.AddCommand("block", command.Command{
		Name:        "add",
		Description: "add a single block.",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			log.Info("block add command")
			return nil
		},
	})
>>>>>>> 3498b7b3eca80b39fc2ab47b0a8596d40667c70e
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
