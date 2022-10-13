package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arifintahu/checkers/testutil/keeper"
	"github.com/arifintahu/checkers/testutil/mock_types"
	"github.com/arifintahu/checkers/x/checkers"
	"github.com/arifintahu/checkers/x/checkers/keeper"
	"github.com/arifintahu/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneGameForRejectGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *mock_types.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := mock_types.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	return server, *k, context, ctrl, bankMock
}

func TestRejectGameByRedOneMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	_, found = keeper.GetStoredGame(ctx, "1")
	require.False(t, found)
}

func TestRejectGameByBlackRefundedGas(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	before := ctx.GasMeter().GasConsumed()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	after := ctx.GasMeter().GasConsumed()
	require.LessOrEqual(t, after, before-5_000)
}
