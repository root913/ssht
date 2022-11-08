package main

import (
	"fmt"
	"os"

	"ssht/client"
	"ssht/config"
	"ssht/ui"
	"ssht/util"

	"github.com/derailed/tview"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	util.Logger.Warn().
		Msg("THIS IS ALPHA VERSION. Passwords are store in configuration file without any encryption.")

	appConfig := config.NewConfig()
	util.Logger.Debug().
		Str("config_path", config.AppHome()).
		Msg("Loading configuration.")

	if err := appConfig.Load(config.AppConfigFile); err != nil {
		util.Logger.Warn().
			Str("config_file_path", config.AppConfigFile).
			Msg("Unable to locate App config. Generating new configuration.")
	}

	appStyles := config.NewStyles()
	appStyles.LoadSkin(appConfig.App.Skin)

	app := &cli.App{}
	app.EnableBashCompletion = true
	app.Name = "SSH Login"
	app.Usage = "SSH Login CLI interface!"
	app.Action = func(c *cli.Context) error {
		if !appConfig.App.HasConnections() {
			cli.ShowAppHelpAndExit(c, 0)
		}
		app := tview.NewApplication()
		ui.NewConnectionsTable(app, appStyles, appConfig)

		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List ssh connections",
			Action: func(c *cli.Context) error {
				connectionTable(appConfig)

				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add new connection",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "password",
					Value: "",
					Usage: "password for connection",
				},
				&cli.StringFlag{
					Name:  "alias",
					Value: "",
					Usage: "alias for connection",
				},
				&cli.StringFlag{
					Name:  "type",
					Value: string(config.PasswordConnection),
					Usage: "type of connection",
				},
				&cli.StringFlag{
					Name:  "key",
					Usage: "Key path for connection",
				},
				&cli.IntFlag{
					Name:  "port",
					Value: 22,
					Usage: "port for connection",
				},
			},
			Action: func(c *cli.Context) error {
				hostArg := c.Args().Get(0)
				if len(hostArg) == 0 {
					util.Logger.Fatal().Msg("Missing host name argument")
					return nil
				}

				connectionType := config.ConnectionType(c.String("type"))
				host, username := getHostAndUsername(hostArg)
				port := c.Int("port")
				alias := c.String("alias")
				password := c.String("password")
				keyPath := c.String("key")

				switch connectionType {
				case config.PasswordConnection:
					if len(password) == 0 {
						password = askPassword()
					}
					conn := config.NewPasswordConnection(host, port, username, password, alias)
					if errs := conn.Validate(); errs != nil {
						util.Logger.Fatal().Err(errs).Msg("")
					}
					appConfig.App.AddConnection(conn)
				case config.KeyConnection:
					if len(keyPath) == 0 {
						keyPath = askKeyPath()
					}
					conn := config.NewKeyConnection(host, port, username, keyPath, alias)
					if errs := conn.Validate(); errs != nil {
						util.Logger.Fatal().Err(errs).Msg("")
					}
					appConfig.App.AddConnection(conn)
				case config.KeyPassphraseConnection:
					if len(keyPath) == 0 {
						keyPath = askKeyPath()
					}
					if len(password) == 0 {
						password = askPassword()
					}
					conn := config.NewKeyPassphraseConnection(host, port, username, keyPath, password, alias)
					if errs := conn.Validate(); errs != nil {
						util.Logger.Fatal().Err(errs).Msg("")
					}
					appConfig.App.AddConnection(conn)
				default:
					if len(password) == 0 {
						password = askPassword()
					}
					conn := config.NewPasswordConnection(host, port, username, password, alias)
					if errs := conn.Validate(); errs != nil {
						util.Logger.Fatal().Err(errs).Msg("")
					}
					appConfig.App.AddConnection(conn)
				}
				appConfig.Save()

				util.Logger.Info().Msg(fmt.Sprintf("Added host %s:%d to config", host, port))
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove connection",
			Action: func(c *cli.Context) error {
				uuid := c.Args().Get(0)
				if len(uuid) == 0 {
					util.Logger.Fatal().Msg("Missing uuid argument")
					return nil
				}
				connection := appConfig.App.Get(uuid)
				if nil == connection {
					util.Logger.Fatal().Msg("Couldn't find connection by given uuid")
					return nil
				}

				appConfig.App.RemoveConnection(connection)
				appConfig.Save()

				util.Logger.Info().Msg("Removed connection from config")
				return nil
			},
		},
		{
			Name:    "connect",
			Aliases: []string{"c"},
			Usage:   "connect to SSH server",
			Action: func(c *cli.Context) error {
				uuidOrAlias := c.Args().Get(0)
				if len(uuidOrAlias) == 0 {
					util.Logger.Fatal().Msg("Missing uuid/alias argument")
					return nil
				}
				connection := appConfig.App.Get(uuidOrAlias)
				if nil == connection {
					util.Logger.Fatal().Msg("Couldn't find connection by given uuid")
					return nil
				}
				client.Connect(connection)

				return nil
			},
		},
		{
			Name:  "alias",
			Usage: "Assign alias to connection",
			Action: func(c *cli.Context) error {
				uuid := c.Args().Get(0)
				alias := c.Args().Get(1)
				if len(uuid) == 0 {
					util.Logger.Fatal().Msg("Missing uuid argument")
					return nil
				}
				if len(alias) == 0 {
					util.Logger.Fatal().Msg("Missing alias argument")
					return nil
				}
				connection := appConfig.App.Get(uuid)
				if nil == connection {
					util.Logger.Fatal().Msg("Couldn't find connection by given uuid")
					return nil
				}
				appConfig.App.SetConnectionAlias(connection.Uuid, alias)

				util.Logger.Info().Msg("Updated alias for connection")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		util.Logger.Fatal().Err(err)
	}
}
