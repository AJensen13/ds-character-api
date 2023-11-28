package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	baseUrl         string
	game            string
	siteMapLocation string
	class           string
}

func main() {
	var sites = []Site{
		{"https://darksouls.fandom.com/wiki", "ds1", "Local_Sitemap", ".mw-allpages-body"},
		// {"", "ds2", ""},
		// {"", "ds3", ""},
		// {"", "er", ""},
		// {"", "bb", ""},
	}

	for _, site := range sites {
		pages := getUrls(site.baseUrl+"/"+site.siteMapLocation, site.class, "")

		for _, page := range pages {
			scrapeSite(page)
			time.Sleep(2 * time.Second)
		}
	}
}

func scrapeSite(url string) {
	// Make an HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Extract data using CSS selectors
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		// Print the text content of each h2 element
		fmt.Println(s.Text())
	})
}

func getUrls(url string, class string, nextLinks string) []string {
	// pass in next links if that's a thing as well as the class
	if nextLinks != "" {

	}

	response := getSiteMap(url)

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var urls []string

	doc.Find(class).Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, link *goquery.Selection) {
            href, exists := link.Attr("href")
            if exists {
                // Print the href value
                fmt.Println(href)
                urls = append(urls, href)
            }
        })
	})

	defer response.Body.Close()

	return urls
}

func getSiteMap(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func parseSiteMap() []string {
	urls := []string{}

	return urls
}
