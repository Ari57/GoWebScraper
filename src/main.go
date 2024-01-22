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
		title, link := ProductTitleLink(e)
		price := ProductPrice(e)
		ProductFormatter(title, link, price)
	})

	// c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone/computers/tablets")
}

func ProductTitleLink(e *colly.HTMLElement) (string, string) {
	title := ""
	link := ""

	e.ForEach(".title", func(_ int, p *colly.HTMLElement) {
		link = e.Request.AbsoluteURL(p.Attr("href"))

		p.Text = strings.Replace(p.Text, "...", "", -1)
		p.Text = strings.Replace(p.Text, "\t", "", -1)
		p.Text = strings.Replace(p.Text, "\n", "", -1)
		title = p.Text
	})
	return title, link
}

func ProductPrice(e *colly.HTMLElement) string {
	price := ""
	e.ForEach(".float-end.price.card-title.pull-right", func(_ int, p *colly.HTMLElement) {
		price = p.Text
	})
	return price
}

func ProductFormatter(title string, link string, price string) {
	fmt.Println(title + " - " + price + " - " + link)
}
