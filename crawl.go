package main

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

func getData(url *url.URL) {

	col := colly.NewCollector()

	col.OnHTML(".amount--3NTpl", func(e *colly.HTMLElement) {
		price := e.Attr("class")
		fmt.Println(e.Text)
		col.Visit(e.Request.AbsoluteURL(price))
	})
	col.OnHTML(".word-break--2nyVq", func(e *colly.HTMLElement) {
		information := e.Attr("class")
		fmt.Println(e.Text)
		col.Visit(e.Request.AbsoluteURL(information))
	})

	col.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL)
	})

	col.Visit(url.String())

}

func main() {
	fmt.Println("Enter Item Type:")
	var item string
	fmt.Scanln(&item)

	fmt.Println("Enter District:")
	var dis string
	fmt.Scanln(&dis)

	///////////////////////////////////////////////

	c := colly.NewCollector()

	c.OnHTML(".gtm-normal-ad a", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("class"))
		title := e.Attr("href")
		fmt.Println(e.Text)
		//fmt.Printf("Title Found: %q -> %s\n", e.Text, title)
		c.Visit(e.Request.AbsoluteURL(title))
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL)
		fmt.Println("===================================")
		url := (r.URL)
		getData(url)
		fmt.Println("===================================")
	})

	c.Visit("https://ikman.lk/en/ads/" + dis + "/" + item)

}
