package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	res := getHttp(baseURL)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find("./pagination").Each()
	fmt.Println(doc)

	return 0
}

func getHttp(sUrl string) *http.Response {
	fmt.Println("[GET]Request :", sUrl)
	req, rErr := http.NewRequest("GET", sUrl, nil)
	checkErr(rErr)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)

	purl, err := url.Parse(sUrl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}
	res, err := client.Do(req)
	checkErr(err)
	checkResCode(res)

	return res
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
