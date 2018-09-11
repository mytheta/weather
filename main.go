package main

import (
	"github.com/mytheta/weather/controller"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "weather"
	app.Usage = "fight the loneliness!"
	app.Action = controller.Weather

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
