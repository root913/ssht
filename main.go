package main

import (
	"github.com/root913/ssht/cmd"
	"github.com/root913/ssht/util"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	util.Logger.Warn().
		Msg("THIS IS ALPHA VERSION. Passwords are store in configuration file without any encryption.")

	cmd.Execute()
}
