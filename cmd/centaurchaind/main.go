package main

import (
	"os"

	"github.com/CentaurDev/centaurchain/app"
	"github.com/CentaurDev/centaurchain/cmd/centaurchaind/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
