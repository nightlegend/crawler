package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/crawler/lib"
	"log"
)

// Extract all http** links from a given webpage
func crawl(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".box").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, status := s.Find("a").Attr("href")
		if status {
			title := s.Find("span").Text()
			fmt.Printf("Review %d: %s - %s\n", i, link, title)
		}

	})
	/*resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return "error"
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns
	log.Println("=====================================")
	bodyByte, _ := ioutil.ReadAll(b)
	resStr := string(bodyByte)
	//log.Println(resStr)
	lib.WriteFile(resStr, url)*/

	return "done"

}

func main() {

	//1. Extract url and download html to local.
	seedUrl := "http://daily.zhihu.com/"
	//seedUrls := []string{"https://docs.docker.com/", "http://daily.zhihu.com/"}
	var strs string
	/*for _, url := range seedUrls {
		strs = crawl(url)
	}*/
	strs = crawl(seedUrl)

	log.Println(strs)
	//2. parse html sample.
	//s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	lib.Parse(strs)

}
