package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mande-chain/x/voting/types"
)

func CmdListAggregateVotesCasted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-aggregate-votes-casted",
		Short: "list all aggregateVotesCasted",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAggregateVotesCastedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AggregateVotesCastedAll(context.Background(), params)
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

func CmdShowAggregateVotesCasted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-aggregate-votes-casted [index]",
		Short: "shows a aggregateVotesCasted",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetAggregateVotesCastedRequest{
				Index: argIndex,
			}

			res, err := queryClient.AggregateVotesCasted(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
