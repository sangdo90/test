package consensus

import "github.com/smartm2m/blockchain/core"

type Consenter interface {
	Mining(*core.Block) error
}
