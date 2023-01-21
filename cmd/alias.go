package cmd

import (
	"fmt"

	"github.com/root913/ssht/config"
	"github.com/root913/ssht/util"

	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:     "alias",
	Aliases: []string{"alias"},
	Short:   "Sets alias for connection",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetConfig()
		uuid := args[0]
		alias := args[1]

		targetConnection := appConfig.App.GetConnection(uuid)
		if nil == targetConnection {
			util.Logger.Fatal().Msg("Couldn't find connection by given uuid")
			return
		}
		aliasConnection := appConfig.App.GetConnection(alias)
		if nil == targetConnection {
			util.Logger.Fatal().Msg(fmt.Sprintf("Given alias is already used by connection %s", aliasConnection.Uuid))
			return
		}

		appConfig.App.SetConnectionAlias(targetConnection.Uuid, alias)
		appConfig.Save()

		util.Logger.Info().Msg("Updated alias for connection")
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
