package cmd

import (
	"fmt"

	"github.com/root913/ssht/config"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func connectionTable(config *config.Config) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Alias", "Host", "Port", "Username", "Key", "Type"})
	for _, connection := range config.App.Connections {
		fmt.Print(connection.Host)
		fmt.Println()
		tw.AppendRow(table.Row{connection.Uuid, connection.Alias, connection.Host, connection.Port, connection.Username, connection.KeyPath, connection.Type})
	}
	tw.SetIndexColumn(1)

	fmt.Println(tw.Render())
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List of connections",
	Run: func(cmd *cobra.Command, args []string) {
		connectionTable(config.GetConfig())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
