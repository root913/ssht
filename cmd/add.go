package cmd

import (
	"fmt"

	"github.com/root913/ssht/config"
	"github.com/root913/ssht/credentials"
	"github.com/root913/ssht/util"

	"github.com/spf13/cobra"
)

var portArg int16 = 22
var connectionTypeArg string = config.PasswordConnection.String()
var passwordArg string = ""
var aliasArg string = ""
var keyPathArg string = ""

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetConfig()

		hostAndUsername := args[0]
		hostArg, usernameArg := util.GetHostAndUsername(hostAndUsername)
		connectionType := config.ConnectionType(connectionTypeArg)
		switch connectionType {
		case config.PasswordConnection:
			if len(passwordArg) == 0 {
				passwordArg = util.AskPassword()
			}
			conn := config.NewPasswordConnection(hostArg, portArg, usernameArg, passwordArg, aliasArg)
			if errs := conn.Validate(); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			if errs := appConfig.App.CheckForduplicates(conn); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			appConfig.App.AddConnection(conn, passwordArg)

			cred := credentials.NewCredentials(config.PassPath)
			cred.Set(conn.Host, config.PasswordConnection.String(), conn.Username, passwordArg)
		case config.KeyConnection:
			if len(keyPathArg) == 0 {
				keyPathArg = util.AskKeyPath()
			}
			conn := config.NewKeyConnection(hostArg, portArg, usernameArg, keyPathArg, aliasArg)
			if errs := conn.Validate(); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			if errs := appConfig.App.CheckForduplicates(conn); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			appConfig.App.AddConnection(conn, "")
		case config.KeyPassphraseConnection:
			if len(keyPathArg) == 0 {
				keyPathArg = util.AskKeyPath()
			}
			if len(passwordArg) == 0 {
				passwordArg = util.AskPassword()
			}
			conn := config.NewKeyPassphraseConnection(hostArg, portArg, usernameArg, keyPathArg, passwordArg, aliasArg)
			if errs := conn.Validate(); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			if errs := appConfig.App.CheckForduplicates(conn); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			appConfig.App.AddConnection(conn, passwordArg)

			cred := credentials.NewCredentials(config.PassPath)
			cred.Set(conn.Host, config.KeyPassphraseConnection.String(), conn.Username, passwordArg)
		default:
			if len(passwordArg) == 0 {
				passwordArg = util.AskPassword()
			}
			conn := config.NewPasswordConnection(hostArg, portArg, usernameArg, passwordArg, aliasArg)
			if errs := conn.Validate(); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			if errs := appConfig.App.CheckForduplicates(conn); errs != nil {
				util.Logger.Fatal().Err(errs).Msg("")
			}
			appConfig.App.AddConnection(conn, passwordArg)

			cred := credentials.NewCredentials(config.PassPath)
			cred.Set(conn.Host, config.PasswordConnection.String(), conn.Username, passwordArg)
		}
		appConfig.Save()

		util.Logger.Info().Msg(fmt.Sprintf("Added host %s:%d to config", hostArg, portArg))
	},
}

func init() {
	addCmd.PersistentFlags().Int16VarP(&portArg, "port", "p", portArg, "Port number")
	addCmd.PersistentFlags().StringVarP(&connectionTypeArg, "connection-type", "t", connectionTypeArg, "Connection type")
	addCmd.PersistentFlags().StringVarP(&passwordArg, "password", "s", "", "Password")
	addCmd.PersistentFlags().StringVarP(&aliasArg, "alias", "a", "", "Alias for connection")
	addCmd.PersistentFlags().StringVarP(&keyPathArg, "key", "k", "", "Key path")
	rootCmd.AddCommand(addCmd)
}
