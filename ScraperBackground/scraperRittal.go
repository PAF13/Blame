package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var artikel []Artikel
var rest int = 0
var id int = 0

type Artikel struct {
	Id            int      `json:"id"`
	URL           string   `json:"url"`
	Artikelnummer []string `json:"artikelnummer"`
}

func Scraper() {
	rand.Seed(time.Now().UnixNano())
	chromedpScrape()
}

func writeJsonFileScrapper(fileName string, data []Artikel) {
	dataJSON, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Blame_"+fileName+".json", dataJSON, 0644)
	if err != nil {
		log.Println(err)
	}
}

func newArtikel(url string) *Artikel {
	return &Artikel{
		URL:           url,
		Artikelnummer: []string{},
	}
}

func chromedpScrape() {
	// Create the initial context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout context
	ctx, cancel = context.WithTimeout(ctx, 99999*time.Second) // Increase timeout duration
	defer cancel()

	// Enable network and request interception
	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		log.Fatalf("Failed to enable network: %v", err)
	}

	// Set up request interception

	// Starting URL
	url := "https://www.rittal.com/de-de/products/PG0900ZUBEHOER1"

	// Track visited URLs to avoid infinite loops
	visited := make(map[string]bool)

	// Start the recursive scraping
	scrapeRittal0(url, &ctx, visited)
	writeJsonFileScrapper("Rittal", artikel)
	fmt.Println("Done")
}

func printRemainingTime(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if !ok {
		log.Println("Context has no deadline")
		return
	}
	remaining := time.Until(deadline)
	log.Printf("Remaining time: %v", remaining)
}

func scrapeRittal0(url string, ctx *context.Context, visited map[string]bool) {
	// Check if the URL has already been visited
	if visited[url] {
		return
	}

	// Mark the URL as visited
	visited[url] = true
	part := newArtikel(url)
	// Log the current URL being visited
	log.Printf("Visiting URL: %s", url)
	printRemainingTime(*ctx)
	// Create a new context for the visit
	visitCtx, visitCancel := context.WithTimeout(*ctx, 99999*time.Second) // Increase individual visit timeout
	defer visitCancel()

	// Variables to capture
	var hrefs []string
	var existsArtikel bool

	err := chromedp.Run(visitCtx,

		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var existsCookie bool
			err := chromedp.EvaluateAsDevTools(`document.querySelector('div.cc-controls button.custom-button.--primary') !== null`, &existsCookie).Do(ctx)
			if err != nil {
				return err
			}
			if existsCookie {
				log.Println("Cookie consent button found, clicking it.")
				return chromedp.Click(`div.cc-controls button.custom-button.--primary`, chromedp.ByQuery).Do(ctx)
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), // Wait for content to load after clicking the button
		chromedp.ActionFunc(func(ctx context.Context) error {
			var existsProductList bool
			err := chromedp.EvaluateAsDevTools(`document.querySelector('div.teaser-grid-wrapper') !== null`, &existsProductList).Do(ctx)
			if err != nil {
				return err
			}

			var existsBestellnummerlist bool
			err = chromedp.EvaluateAsDevTools(`document.querySelector('div.content-table') !== null`, &existsBestellnummerlist).Do(ctx)
			if err != nil {
				return err
			}

			err = chromedp.EvaluateAsDevTools(`document.querySelector('div.breadcrumb-wrapper.--active') !== null`, &existsArtikel).Do(ctx)
			if err != nil {
				return err
			}

			if existsProductList {
				log.Println("On Product list page")
				err := chromedp.Evaluate(`Array.from(document.querySelectorAll('div.teaser-grid-wrapper a.custom-link')).map(a => a.href)`, &hrefs).Do(ctx)
				if err != nil {
					return err
				}
			} else if existsBestellnummerlist {
				log.Println("On Bestell list page")
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('div.content-table a')).map(a => a.href)`, &hrefs).Do(ctx)
				if err != nil {
					return err
				}
				rest = rest + len(hrefs)
			} else if existsArtikel {
				log.Println("On Artikel page")
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('div.product-info span.value')).map(span => span.textContent)`, &part.Artikelnummer).Do(ctx)
				if err != nil {
					return err
				}
				id++
				part.Id = id
				artikel = append(artikel, *part)
				fmt.Printf("num: %-20d", part.Id)
				fmt.Printf("Artikelnummer: %-40s", part.Artikelnummer)
				fmt.Printf("URL: %-50s", part.URL)
				fmt.Printf("\n")
			} else {
				log.Println("Neither Product list nor Bestell list found on the page.")
			}
			return nil
		}),
	)
	if err != nil {
		log.Printf("Failed to visit %s: %v", url, err)
		return
	}

	fmt.Println("Progress: " + fmt.Sprintf("%d", id) + " / " + fmt.Sprintf("%d", rest))
	// Check if any hrefs were found
	if len(hrefs) == 0 {
		log.Println("No links found on the page.")
		return
	}

	// Print the captured hrefs and recursively visit them
	for _, href := range hrefs {
		log.Printf("Found link: %s", href)
		// Recursively scrape the linked page
		scrapeRittal0(href, ctx, visited)
	}
}
