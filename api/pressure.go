package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assemblaj/zapata_takehome/utils"
)

type PressureHandler struct{}

func (p PressureHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	res, err := utils.GetPageData(rw, r)
	if err != nil {
		log.Fatal(err)
	} else {
		valMap := utils.GetKeyValueFromString(res)
		pressure, hasPressure := valMap["Pressure"]
		if hasPressure {
			fmt.Fprintf(rw, pressure)
		} else {
			fmt.Fprintf(rw, "Location has no pressure")
		}
	}
}
