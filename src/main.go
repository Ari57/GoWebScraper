package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "computer") || strings.Contains(link, "phone") {
			fmt.Println(e.Request.AbsoluteURL(link))
		}
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("")
	// 	// fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
}
