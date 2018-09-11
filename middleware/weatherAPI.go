package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/mytheta/weather/model"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWeather(region string) *model.Weather {

	API_KEY := os.Getenv("WEATHER_API")
	BASE_URL := "http://api.openweathermap.org/data/2.5/weather"

	resp, err := http.Get(BASE_URL + "?q=" + region + ",jp&APPID=" + API_KEY)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(model.Weather)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	data.Main.Temp -= 273.15

	return data
}
