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
	console.RegisterBlockchain(bc, core.NewBlockchain)
	console.RegisterBlock(core.NewBlock)
	RegisterCommand()
	console.Start()
}

// RegisterCommand register commands for manage blockchains.
func RegisterCommand() {
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
