package cli

import (
	"strconv"
	"context"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mun/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdModuleAccountBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module-account-balance",
		Short: "Query the current claim module's account balance",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryModuleAccountBalanceRequest{}
			res, err := queryClient.ModuleAccountBalance(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
