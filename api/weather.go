package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assemblaj/zapata_takehome/utils"
)

type WeatherHandler struct{}

func (w WeatherHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	res, err := utils.GetPageData(rw, r)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(rw, res)
	}
}
