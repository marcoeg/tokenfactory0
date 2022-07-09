package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/marcog/tokenfactory/testutil/keeper"
	"github.com/marcog/tokenfactory/x/tokenfactory/keeper"
	"github.com/marcog/tokenfactory/x/tokenfactory/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TokenfactoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
