package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkResCode(res)
	goquery.ToEnd
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkResCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
