package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	title := ""
	price := ""

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "computer") {
			c.Visit(e.Request.AbsoluteURL(link))

			c.OnHTML(".wrapper", func(e *colly.HTMLElement) {
				e.ForEach("h4", func(_ int, p *colly.HTMLElement) {
					p.Text = strings.Replace(p.Text, "\t", "", -1)
					p.Text = strings.Replace(p.Text, "\n", "", -1)

					if strings.Contains(p.Text, "$") {
						price = p.Text
					} else {
						p.Text = strings.Replace(p.Text, "...", "", -1)
						title = p.Text
					}

					product := title + " - " + price
					fmt.Println(product)
				})

			})
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("")
	// 	// fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
}
