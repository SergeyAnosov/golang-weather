package main

import (
	"flag"
	"fmt"
	"golang/weather/geo"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println("city: ", *city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
}
