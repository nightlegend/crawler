package lib

import (
	"fmt"
	"strings"
	"log"
	"golang.org/x/net/html"
)


func F(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		F(c)
	}
}

func Parse(str string)  {
	doc, err := html.Parse(strings.NewReader(str))
	if err != nil {
		log.Fatal(err)
	}
	F(doc)
}

