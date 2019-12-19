package main

import (
	"fmt"
	"os"
	"runtime"

	cmd "github.com/SebastianJ/elrond-cli/cmd/subcommands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Elrond CLI"
	app.Version = fmt.Sprintf("%s/%s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	app.Usage = "Interact with Elrond's blockchain using CLI commands"

	app.Authors = []cli.Author{
		{
			Name:  "Sebastian Johnsson",
			Email: "",
		},
	}

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, cmd.TransferCommand())
	app.Commands = append(app.Commands, cmd.BalanceCommand())

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api-endpoint",
			Usage: "Which API endpoint to use for API commands",
			Value: "https://wallet-api.elrond.com",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}
}
