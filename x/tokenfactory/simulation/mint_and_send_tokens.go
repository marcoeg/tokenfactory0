package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/marcoeg/tokenfactory/x/tokenfactory/keeper"
	"github.com/marcoeg/tokenfactory/x/tokenfactory/types"
)

func SimulateMsgMintAndSendTokens(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintAndSendTokens{
			Owner: simAccount.Address.String(),
		}

		// TODO: Handling the MintAndSendTokens simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MintAndSendTokens simulation not implemented"), nil, nil
	}
}
