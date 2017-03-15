package lib

import (
	"os"
	"fmt"
	"log"
)

func check(e error)  {
	if e!= nil {
		panic(e)
	}
}

func WriteFile(str string, url string)  {
	log.Println("the url is "+url)
	f, err := os.Create("/Users/davidguo/Desktop/crawlerdata/"+url[6:15]+".txt")
	check(err)
	defer f.Close()
	n3, err := f.WriteString(str)
	fmt.Printf("wrote %d bytes\n", n3)
}