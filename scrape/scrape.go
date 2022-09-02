package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetData(e *colly.XMLElement) {
	fmt.Println("Found data!")
	fmt.Println(e.ChildText("title"))
	fmt.Println(e.ChildText("yt:videoId"))
}

func RunScrape() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnXML("//feed/entry", GetData)
	// c.OnXML("entry", GetData)
	c.Visit("https://www.youtube.com/feeds/videos.xml?channel_id=UCgA2jKRkqpY_8eysPUs8sjw")

}
