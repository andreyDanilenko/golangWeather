package main

import (
	"app/weather/geo"
	"app/weather/weather"
	"flag"
	"fmt"
	"io"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)

	if err == io.EOF {
		fmt.Println(err.Error())
	}

	weather := weather.GetWeather(*geoData, *format)

	fmt.Println(*geoData)
	fmt.Println(weather)
}
