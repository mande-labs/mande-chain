package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mande-chain/x/voting/types"
)

func CmdListAggregateVotesReceived() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-aggregate-votes-received",
		Short: "list all aggregateVotesReceived",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAggregateVotesReceivedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AggregateVotesReceivedAll(context.Background(), params)
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

func CmdShowAggregateVotesReceived() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-aggregate-votes-received [index]",
		Short: "shows a aggregateVotesReceived",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetAggregateVotesReceivedRequest{
				Index: argIndex,
			}

			res, err := queryClient.AggregateVotesReceived(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
