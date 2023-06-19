package eth

import (
	"context"

	"github.com/gatechain/go-ethereum/core"
	"github.com/gatechain/go-ethereum/core/state"
	"github.com/gatechain/go-ethereum/core/types"
	"github.com/gatechain/go-ethereum/core/vm"
	"github.com/gatechain/go-ethereum/eth/tracers"
	"github.com/gatechain/go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
