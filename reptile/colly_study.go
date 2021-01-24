package reptile

//爬虫

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func Start() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	c.AllowedDomains = []string{"studygolang.com"}

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		if strings.Contains(e.Text, "go") {
			fmt.Printf("Link found: %q -> %s\n", e.Text, e.Request.AbsoluteURL(link))
		}
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// Before making a request print "Visiting ..."
	//在请求之前打印请求路径
	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on https://studygolang.com
	//开始爬虫
	c.Visit("https://studygolang.com")

}
