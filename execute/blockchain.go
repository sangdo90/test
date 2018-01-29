package execute

import (
	"errors"
	"fmt"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

//BlockchainCommands is ...
func BlockchainCommands() {
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
		Description: "manage a single block.",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			log.Info("block command")
			return nil
		},
	})
	command.AddCommand("", command.Command{
		Name:        "print",
		Description: "print infomation.",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			log.Info("print command")
			return nil
		},
	})
	command.AddCommand("print", command.Command{
		Name:        "blockchain",
		Description: "print a blockchain.",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			fmt.Println(console.GetContext().Blockchain.String())
			return nil
		},
	})
	command.AddCommand("", command.Command{
		Name:        "chain",
		Description: "manage a block chain",
		Commands:    make(command.CommandSlice, 0),
		Flags:       make(command.FlagSlice, 0),
		Run: func(args []string) error {
			log.Info("chain command")
			return nil
		},
	})
	command.AddCommand("", command.Command{
		Name:        "quit",
		Description: "Exit the program",
		Commands:    nil,
		Flags:       nil,
		Run: func(args []string) error {
			//this line does not work on windows 10, need to modify it.
			//syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			return nil
		},
	})
}

//CreateBlockChain creates a blockchain containing genesisblock
func CreateBlockChain() error {
	bc := core.NewBlockChain()
	bc.RegisterBlockChain()

	return nil
}

//NewBlockToBlockChain appends 'blk block' to the 'bidx index' of the 'bcid blockchain'.
func NewBlockToBlockChain(bcid uint64, bidx uint64, blk console.Blocker) error {

	bc, err := core.SelectBlockChain(bcid)
	if err != nil {
		return err
	}

	if bidx > bc.BlockChainHeight-1 {
		return errors.New("Invalid block index, Genesis block index is zero(0)")
	}

	if bidx == bc.BlockChainHeight-1 {
		bc.AddBlock(blk)
	}

	if bidx < bc.BlockChainHeight-1 {
		cbc := core.CutBlockChain(*bc, bidx)
		cbc.RegisterBlockChain()
		cbc.AddBlock(blk)
	}
	return nil
}
