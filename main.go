package main

import (
	"log"
	"net/http"

	"github.com/assemblaj/zapata_takehome/api"
)

func main() {
	http.Handle("/weather/", api.WeatherHandler{})
	http.Handle("/temperature/", api.TemperatureHandler{})
	http.Handle("/pressure/", api.PressureHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
