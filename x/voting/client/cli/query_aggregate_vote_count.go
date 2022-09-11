package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mande-chain/x/voting/types"
)

func CmdListAggregateVoteCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-aggregate-vote-count",
		Short: "list all aggregateVoteCount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAggregateVoteCountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AggregateVoteCountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAggregateVoteCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-aggregate-vote-count [index]",
		Short: "shows a aggregateVoteCount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetAggregateVoteCountRequest{
				Index: argIndex,
			}

			res, err := queryClient.AggregateVoteCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
