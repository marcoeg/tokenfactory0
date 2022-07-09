package keeper

import (
	"github.com/marcog/tokenfactory/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
