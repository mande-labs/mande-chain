package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"mande-chain/x/voting/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group voting queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListVoteBook())
	cmd.AddCommand(CmdShowVoteBook())
	cmd.AddCommand(CmdListAggregateVotesCasted())
	cmd.AddCommand(CmdShowAggregateVotesCasted())
	cmd.AddCommand(CmdListAggregateVotesReceived())
	cmd.AddCommand(CmdShowAggregateVotesReceived())
	// this line is used by starport scaffolding # 1

	return cmd
}
