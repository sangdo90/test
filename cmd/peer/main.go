package main

import (
	"github.com/smartm2m/blockchain/execute"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

func main() {
	log.SetLogLevel(log.DebugLogLevel)
	RegisterCommand()
	console.Start()
}

// RegisterCommand register commands for manage blockchains.
func RegisterCommand() {
	execute.BlockchainCommands()
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
