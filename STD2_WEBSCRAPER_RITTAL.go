package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func scraper(url []string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 99999*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		log.Fatalf("Failed to enable network: %v", err)
	}

	for _, b := range url {
		scrapeRittal0(b, &ctx)
	}

	fmt.Println("Done")
}

type ArtikelSeedRaw struct {
	URL           string   `json:"url"`
	Artikelnummer []string `json:"artikelnummer"`
	Bestellnummer []string `json:"bestellnummer"`
}
type ArtikelSeed struct {
	URL           string `json:"url"`
	Artikelnummer string `json:"artikelnummer"`
	Bestellnummer string `json:"bestellnummer"`
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

func scrapeRittal0(url string, ctx *context.Context) {
	log.Printf("Visiting URL: %s", url)
	printRemainingTime(*ctx)

	visitCtx, visitCancel := context.WithTimeout(*ctx, 99999*time.Second)
	defer visitCancel()

	part := newArtikel(url)
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
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {

			err := chromedp.EvaluateAsDevTools(`document.querySelector('div.breadcrumb-wrapper.--active') !== null`, &existsArtikel).Do(ctx)
			if err != nil {
				return err
			}

			if existsArtikel {
				log.Println("On Artikel page")
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('div.product-info span.value')).map(span => span.textContent)`, &part.Artikelnummer).Do(ctx)
				if err != nil {
					return err
				}

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

}
