package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arifintahu/checkers/testutil/keeper"
	"github.com/arifintahu/checkers/x/checkers"
	"github.com/arifintahu/checkers/x/checkers/keeper"
	"github.com/arifintahu/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
