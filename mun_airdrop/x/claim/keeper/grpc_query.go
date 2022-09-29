package keeper

import (
	"mun/x/claim/types"
)

var _ types.QueryServer = Keeper{}
