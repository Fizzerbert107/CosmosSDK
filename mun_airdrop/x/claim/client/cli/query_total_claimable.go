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

func CmdTotalClaimable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-claimable [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the total claimable amount remaining for an account.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the total claimable amount remaining for an account.
Example:
$ %s query claim total-claimable osmo1ey69r37gfxvxg62sh4r0ktpuc46pzjrm23kcrx
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
			// Query store
			res, err := queryClient.TotalClaimable(context.Background(), &types.QueryTotalClaimableRequest{
				Address: args[0],
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
