package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gocolly/colly"
)

var (
	titleSlice       []string
	descriptionSlice []string
	linkSlice        []string
	priceSlice       []string
)

func main() {
	file := CreateDocument()

	c := colly.NewCollector()
	c.AllowURLRevisit = false

	done := make(chan struct{})

	go func() {
		FindElements(c)
		close(done)
	}()

	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")

	<-done

	WriteData(file, titleSlice, descriptionSlice, priceSlice, linkSlice)
}

func FindElements(c *colly.Collector) {
	ProductPages := []string{}

	c.OnHTML(".nav.flex-column", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, p *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(p.Attr("href"))

			if !slices.Contains(ProductPages, link) {
				ProductPages = append(ProductPages, link)
				fmt.Println(" ")
				fmt.Println("Visiting", link)
				fmt.Println(" ")
				c.Visit(link)

				c.OnHTML(".caption", func(e *colly.HTMLElement) {
					title, link := ProductTitleLink(e)
					description := ProductDescription(e)
					price := ProductPrice(e)

					titleSlice, descriptionSlice, linkSlice, priceSlice = ProductFormatter(title, description, link, price)

				})
			}
		})
	})
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

func ProductDescription(e *colly.HTMLElement) string {
	description := ""

	e.ForEach(".description.card-text", func(_ int, p *colly.HTMLElement) {
		description = p.Text
	})

	return description

}

func ProductPrice(e *colly.HTMLElement) string {
	price := ""
	e.ForEach(".float-end.price.card-title.pull-right", func(_ int, p *colly.HTMLElement) {
		price = p.Text
	})
	return price
}

func ProductFormatter(title string, description string, link string, price string) ([]string, []string, []string, []string) {
	titleSlice = append(titleSlice, title)
	descriptionSlice = append(descriptionSlice, description)
	linkSlice = append(linkSlice, link)
	priceSlice = append(priceSlice, price)

	return titleSlice, descriptionSlice, linkSlice, priceSlice
}
