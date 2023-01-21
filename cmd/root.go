package cmd

import (
	"os"

	"github.com/root913/ssht/config"
	"github.com/root913/ssht/ui"
	"github.com/root913/ssht/util"
	"github.com/rs/zerolog"

	"github.com/derailed/tview"
	"github.com/spf13/cobra"
)

const DefaultLogLevel string = "debug"

var LogLevel string

var rootCmd = &cobra.Command{
	Use:     "ssht",
	Short:   "ssht - a simple CLI to managing your SSH connections",
	Version: "1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		zerolog.SetGlobalLevel(parseLevel(LogLevel))
	},
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

func init() {
	rootCmd.Flags().StringVarP(
		&LogLevel,
		"logLevel", "l",
		DefaultLogLevel,
		"Specify a log level (info, warn, debug, trace, error)",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.Logger.Err(err).Msg("Whoops. There was an error while executing your CLI")
		os.Exit(1)
	}
}

func parseLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
