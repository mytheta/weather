package controller

import (
	"fmt"
	"github.com/mytheta/weather/middleware"
	"github.com/urfave/cli"
	"strconv"
)

func Weather(context *cli.Context) error {

	region := context.Args().Get(0)

	data := middleware.GetWeather(region)
	//stringに変換
	humidity := strconv.Itoa(data.Main.Humidity)

	temp := strconv.FormatFloat(data.Main.Temp, 'f', 4, 64)


	fmt.Println(region + "の天気は" + data.Weather[0].Main + ", 気温:"+ temp + "度, 湿度:" + humidity + "%")
	return nil
}