package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mande-chain/x/voting/types"
)

var _ = strconv.Itoa(0)

func CmdCredibilityScore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credibility-score [address]",
		Short: "Query credibility-score",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCredibilityScoreRequest{

				Address: reqAddress,
			}

			res, err := queryClient.CredibilityScore(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
