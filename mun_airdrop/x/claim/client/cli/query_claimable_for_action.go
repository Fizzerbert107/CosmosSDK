package cli

import (
	"strconv"
	"context"
	"fmt"
	"strings"
	"github.com/cosmos/cosmos-sdk/version"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mun/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdClaimableForAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claimable-for-action [address] [action]",
		Args:  cobra.ExactArgs(2),
		Short: "Query an address' claimable amount for a specific action",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query an address' claimable amount for a specific action

Example:
$ %s query claim claimable-for-action osmo1ey69r37gfxvxg62sh4r0ktpuc46pzjrm23kcrx ActionAddLiquidity
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			action, ok := types.Action_value[args[1]]
			if !ok {
				return fmt.Errorf("invalid action type: %s", args[1])
			}

			// Query store
			res, err := queryClient.ClaimableForAction(context.Background(), &types.QueryClaimableForActionRequest{
				Address: args[0],
				Action:  types.Action(action),
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintObjectLegacy(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
