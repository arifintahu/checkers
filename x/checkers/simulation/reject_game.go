package simulation

import (
	"math/rand"

	"github.com/arifintahu/checkers/x/checkers/keeper"
	"github.com/arifintahu/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRejectGame(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRejectGame{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RejectGame simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RejectGame simulation not implemented"), nil, nil
	}
}
