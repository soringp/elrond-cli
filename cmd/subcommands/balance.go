package cmd

import (
	"errors"
	"fmt"

	"github.com/SebastianJ/elrond-cli/api"
	"github.com/SebastianJ/elrond-cli/utils"
	"github.com/urfave/cli"
)

// BalanceCommand sets up the CLI functionality for performing balance checks
func BalanceCommand() cli.Command {
	return cli.Command{
		Name:  "balance",
		Usage: "check the balance of a specific address",
		Action: func(ctx *cli.Context) error {
			return sendBalanceCommand(ctx)
		},
	}
}

func sendBalanceCommand(ctx *cli.Context) error {
	apiHost := ctx.GlobalString("api-endpoint")
	address := ctx.Args().Get(0)

	if address == "" || len(address) < 64 {
		return errors.New("please provide a valid address")
	}

	accountData, err := api.GetBalance(address, apiHost)

	if err != nil {
		return errors.New("failed to retrieve balance")
	}

	balance := accountData.Balance
	val, _ := utils.ConvertNumeralStringToBigFloat(balance)

	fmt.Println(fmt.Sprintf("Balance for %s is: %f", address, val))

	return nil
}
