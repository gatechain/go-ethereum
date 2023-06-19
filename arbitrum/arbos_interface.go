package arbitrum

import (
	"context"

	"github.com/gatechain/go-ethereum/arbitrum_types"
	"github.com/gatechain/go-ethereum/core"
	"github.com/gatechain/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
