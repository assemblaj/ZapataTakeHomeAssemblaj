package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assemblaj/zapata_takehome/utils"
)

type WeatherHandler struct{}

func (w WeatherHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	/*
		fileName := path.Base(r.URL.Path)
		//api(rootPage + fileName)
		url := utils.RootPage + fileName
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(rw, "404 Page not found")
		} else {
			defer response.Body.Close()
			body, err2 := io.ReadAll(response.Body)
			if err2 != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(rw, string(body))
		} */
	res, err := utils.GetPageData(rw, r)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(rw, res)
	}
}
