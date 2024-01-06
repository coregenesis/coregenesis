package root

import (
	"fmt"
	"os"

	"github.com/coregenesis/coregenesis/command/backup"
	"github.com/coregenesis/coregenesis/command/genesis"
	"github.com/coregenesis/coregenesis/command/helper"
	"github.com/coregenesis/coregenesis/command/ibft"
	"github.com/coregenesis/coregenesis/command/license"
	"github.com/coregenesis/coregenesis/command/monitor"
	"github.com/coregenesis/coregenesis/command/peers"
	"github.com/coregenesis/coregenesis/command/secrets"
	"github.com/coregenesis/coregenesis/command/server"
	"github.com/coregenesis/coregenesis/command/status"
	"github.com/coregenesis/coregenesis/command/txpool"
	"github.com/coregenesis/coregenesis/command/version"
	"github.com/coregenesis/coregenesis/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Genesis Core is an Ethereum-compatible Blockchain network",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
