package controller

import (
	"fmt"
	"github.com/urfave/cli"
)

func Weather(context *cli.Context) error {
	fmt.Println(context.Args().Get(0))
	return nil
}