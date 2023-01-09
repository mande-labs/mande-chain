package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mande-chain/x/oracle/types"
)

// CmdNetworkConstantResult queries request result by reqID
func CmdNetworkConstantResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-constant-result [request-id]",
		Short: "Query the NetworkConstant result data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.NetworkConstantResult(context.Background(), &types.QueryNetworkConstantRequest{RequestId: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdLastNetworkConstantID queries latest request
func CmdLastNetworkConstantID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last-network-constant-id",
		Short: "Query the last request id returned by NetworkConstant ack packet",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.LastNetworkConstantId(context.Background(), &types.QueryLastNetworkConstantIdRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
