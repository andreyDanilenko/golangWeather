package weather

import (
	"app/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
