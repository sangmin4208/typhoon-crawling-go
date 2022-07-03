package api

import (
	"log"
	"net/http"
)

func GetTyphoon() *http.Response {
	res, err := http.Get("https://www.weather.go.kr/w/typhoon/report.do#")
	if err != nil {
		log.Fatal(err)
	}
	return res
}