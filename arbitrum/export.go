package arbitrum

import (
	"context"

	"github.com/gatechain/go-ethereum/common/hexutil"
	"github.com/gatechain/go-ethereum/core"
	"github.com/gatechain/go-ethereum/internal/ethapi"
	"github.com/gatechain/go-ethereum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
