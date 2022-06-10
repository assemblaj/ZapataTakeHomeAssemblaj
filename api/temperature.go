package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assemblaj/zapata_takehome/utils"
)

type TemperatureHandler struct{}

func (t TemperatureHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	res, err := utils.GetPageData(rw, r)
	if err != nil {
		log.Fatal(err)
	} else {
		valMap := utils.GetKeyValueFromString(res)
		temperature, hasTemp := valMap["Temperature"]
		if hasTemp {
			fmt.Fprintf(rw, temperature)
		} else {
			fmt.Fprintf(rw, "Location has no Temperature")
		}
	}
}
