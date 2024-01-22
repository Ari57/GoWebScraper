package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// link := e.Attr("href")
	// c.Visit(e.Request.AbsoluteURL(link))

	c.OnHTML(".caption", func(e *colly.HTMLElement) {
		ProductTitle(e)
		ProductPrice(e)
	})

	// c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone/computers/tablets")
}

func ProductTitle(e *colly.HTMLElement) {
	e.ForEach(".title", func(_ int, p *colly.HTMLElement) {

		p.Text = strings.Replace(p.Text, "...", "", -1)
		p.Text = strings.Replace(p.Text, "\t", "", -1)
		p.Text = strings.Replace(p.Text, "\n", "", -1)

		fmt.Print(p.Text + " - ")

	})
}

func ProductPrice(e *colly.HTMLElement) {
	e.ForEach(".float-end.price.card-title.pull-right", func(_ int, p *colly.HTMLElement) {
		fmt.Print(p.Text)
		fmt.Println()
	})
}
