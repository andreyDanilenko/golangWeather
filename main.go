package main

import (
	"app/weather/geo"
	"flag"
	"fmt"
	"io"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()
	// fmt.Println(*city)
	// fmt.Println(*format)

	// read := strings.NewReader("Привет!")
	// block := make([]byte, 4)

	geoData, err := geo.GetMyLocation(*city)

	if err == io.EOF {
		fmt.Println(err.Error())
	}

	fmt.Println(*geoData)

	// for {
	// 	_, err := read.Read(block)
	// 	fmt.Printf("%q\n", block)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

}
