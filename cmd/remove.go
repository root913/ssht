package cmd

import (
	"fmt"

	"github.com/root913/ssht/config"
	"github.com/root913/ssht/util"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Removes connection",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetConfig()
		uuidOrAliasArg := args[0]
		connection := appConfig.App.Get(uuidOrAliasArg)
		if nil == connection {
			util.Logger.Fatal().Msg("Couldn't find connection by given uuid/alias")
			return
		}

		appConfig.App.RemoveConnection(connection)
		appConfig.Save()

		util.Logger.Info().
			Msg(fmt.Sprintf("Removed connection %s from config", uuidOrAliasArg))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
