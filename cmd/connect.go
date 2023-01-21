package cmd

import (
	"github.com/root913/ssht/client"
	"github.com/root913/ssht/config"
	"github.com/root913/ssht/util"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to SSH server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetConfig()
		uuidOrAliasArg := args[0]
		connection := appConfig.App.GetConnection(uuidOrAliasArg)
		if nil == connection {
			util.Logger.Fatal().Msg("Couldn't find connection by given uuid/alias")
			return
		}

		client.Connect(connection)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
