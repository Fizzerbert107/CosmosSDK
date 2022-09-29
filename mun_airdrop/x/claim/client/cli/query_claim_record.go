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

func CmdClaimRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-record [address]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the claim record for an account.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the claim record for an account.
This contains an address' initial claimable amounts, and the completed actions.

Example:
$ %s query claim claim-record <address>
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
			res, err := queryClient.ClaimRecord(context.Background(), &types.QueryClaimRecordRequest{Address: args[0]})
			if err != nil {
				return err
			}
			return clientCtx.PrintObjectLegacy(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
