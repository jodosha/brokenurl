package main

import (
	"fmt"

	"github.com/asciimoo/colly"
	"github.com/fatih/color"
)

func main() {
	url := "http://hanamirb.org"

	okPrinter := color.New(color.FgGreen)
	errorPrinter := color.New(color.FgRed)

	// Instantiate default collector
	c := colly.NewCollector()

	// Visit only hanamirb.org
	c.AllowedDomains = []string{"hanamirb.org"}

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*hanamirb.*",
		Parallelism: 2,
		// Delay:       5 * time.Second,
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {

		okPrinter.Println(" ok")
	})

	c.OnError(func(r *colly.Response, err error) {
		errorPrinter.Println(" ", err)
	})
	c.Visit(url)
	c.Wait()
}
