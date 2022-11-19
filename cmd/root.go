package cmd

import (
	"os"

	"github.com/root913/ssht/config"
	"github.com/root913/ssht/ui"
	"github.com/root913/ssht/util"

	"github.com/derailed/tview"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssht",
	Short: "ssht - a simple CLI to managing your SSH connections",
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetConfig()
		appStyles := config.GetStyles()
		if !appConfig.App.HasConnections() {
			cmd.Help()
			os.Exit(0)
		}
		app := tview.NewApplication()
		ui.NewConnectionsTable(app, appStyles, appConfig)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.Logger.Err(err).Msg("Whoops. There was an error while executing your CLI")
		os.Exit(1)
	}
}
