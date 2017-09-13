package lib

import (
	"golang.org/x/net/html"
	"log"
	"strings"
	"fmt"
)

func parseContent(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				//fmt.Println(a.Val)
				break
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "box" {
				fmt.Println(n.DataAtom)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseContent(c)
	}
}

func Parse(str string) {
	doc, err := html.Parse(strings.NewReader(str))
	if err != nil {
		log.Fatal(err)
	}
	parseContent(doc)
}
