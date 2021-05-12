package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		// http://www.chinagwy.org/html/stzx/xinjiang/index.html
		colly.AllowedDomains("www.chinagwy.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		m := map[string]int32{}
		if strings.HasSuffix(link, ".pdf") || strings.Contains(e.Text, "新疆") {
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
		}
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.chinagwy.org/html/stzx/xinjiang/index.html")
}
