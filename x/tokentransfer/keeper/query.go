package keeper

import (
	"tokenfactory/x/tokentransfer/types"
)

var _ types.QueryServer = Keeper{}
