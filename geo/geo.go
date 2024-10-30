package geo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("There isn't city")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	res, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)

	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	res, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return false
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return false
	}

	var responsePopulation CityPopulationResponse
	json.Unmarshal(body, &responsePopulation)
	return !responsePopulation.Error
}
