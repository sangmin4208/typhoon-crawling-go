package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/sangmin4208/typhoon-crawling-go/typhoon/api"
	"github.com/sangmin4208/typhoon-crawling-go/typhoon/parse"
	"github.com/sangmin4208/typhoon-crawling-go/typhoon/write"
)

func main() {
	res := api.GetTyphoon()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	title := parse.ParseTyphoonInfo(bytes.NewReader(body))
	list := parse.ParseTyphoonTable(bytes.NewReader(body))
	defer res.Body.Close()
	write.TyphoonList("output", title, list)
}
