package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mun/x/claim/types"
)
var _ types.QueryServer = Keeper{}

func (k Keeper) ModuleAccountBalance(c context.Context, _ *types.QueryModuleAccountBalanceRequest) (*types.QueryModuleAccountBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	moduleAccBal := sdk.NewCoins(k.GetModuleAccountBalance(ctx))

	return &types.QueryModuleAccountBalanceResponse{ModuleAccountBalance: moduleAccBal}, nil
}