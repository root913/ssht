package main

import (
	"github.com/root913/ssht/cmd"
	"github.com/root913/ssht/util"
)

func main() {
	util.Logger.Warn().
		Msg("THIS IS ALPHA VERSION.")

	cmd.Execute()
}
