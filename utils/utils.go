package utils

import (
	"io"
	"net/http"
	"path"
	"regexp"
)

const RootPage string = "https://tgftp.nws.noaa.gov/data/observations/metar/decoded/"

func GetPageData(rw http.ResponseWriter, r *http.Request) (string, error) {
	fileName := path.Base(r.URL.Path)
	url := RootPage + fileName
	response, err := http.Get(url)
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()
		body, err2 := io.ReadAll(response.Body)
		if err2 != nil {
			return "", err2
		}
		return string(body), nil
	}

}

func GetKeyValueFromString(str string) map[string]string {
	var rex = regexp.MustCompile("(\\w+): (\\w+)")
	data := rex.FindAllStringSubmatch(str, -1)
	res := make(map[string]string)
	for _, kv := range data {
		k := kv[1]
		v := kv[2]
		res[k] = v
	}
	return res
}
