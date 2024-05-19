package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type ArtikelSeed struct {
	Bestellnummer string
	Artikelnummer string
	URL           string
}

/*
Seed based data retrievel:
this code will get part data such as weight, tech. data etc. with only Bestell- or Artikelnummer and the set algorithm
*/
func main() {
	//importJSON()
	//Scraper()
	hersteller := "SIEMENS"
	switch hersteller {
	case "SIEMENS":
		scraperSiemens()
	}

	fmt.Println("Done")
}

type Scrapelocation struct {
	Hersteller string
	location   int
}

func scraperSiemens() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 99999*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		log.Fatalf("Failed to enable network: %v", err)
	}

	scrapelocation := []Scrapelocation{
		{Hersteller: "weidmueller",
			location: 335},
		{Hersteller: "schneider-electric",
			location: 335},
		{Hersteller: "danfoss",
			location: 23},
		{Hersteller: "eaton",
			location: 335},
		{Hersteller: "festo",
			location: 335},
		{Hersteller: "helukabel",
			location: 335},
		{Hersteller: "ifm-electronic",
			location: 329},
		{Hersteller: "lapp",
			location: 335},
		{Hersteller: "murrelektronik",
			location: 335},
		{Hersteller: "pepperl-fuchs",
			location: 335},
		{Hersteller: "pilz",
			location: 45},
		{Hersteller: "woehner",
			location: 1},
		{Hersteller: "sick",
			location: 335},
		{Hersteller: "d-link",
			location: 22},
	}

	for _, b := range scrapelocation {
		url := "https://www.conrad.de/de/marken/" + b.Hersteller + ".html"
		page := 2
		var produkte map[string]*Product = make(map[string]*Product)
		scrapesSiemensLoop(url, &ctx, produkte)
		for page < b.location {
			scrapesSiemensLoop(url+"?page="+fmt.Sprintf("%d", page), &ctx, produkte)
			page++
		}
		writeJsonFileScrapper(b.Hersteller, produkte)
	}
}
func scrapesSiemensLoop(url string, ctx *context.Context, produkte map[string]*Product) {
	product := newProduct(url)

	log.Printf("Visiting URL: %s", url)
	printRemainingTime(*ctx)

	visitCtx, visitCancel := context.WithTimeout(*ctx, 240*time.Second)
	defer visitCancel()

	var existsArtikel bool
	err := chromedp.Run(visitCtx,

		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var existsCookie bool
			err := chromedp.EvaluateAsDevTools(`document.querySelector('button.sc-dcJsry ghqBLM') !== null`, &existsCookie).Do(ctx)
			if err != nil {
				return err
			}
			if existsCookie {
				log.Println("Cookie consent button found, clicking it.")
				return chromedp.Click(`button.sc-dcJsry ghqBLM`, chromedp.ByQuery).Do(ctx)
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), // Wait for content to load after clicking the button
		chromedp.ActionFunc(func(ctx context.Context) error {

			err := chromedp.EvaluateAsDevTools(`document.querySelector('p.product__manufactorId') !== null`, &existsArtikel).Do(ctx)
			if err != nil {
				return err
			}
			var bestellnummer []string
			var kurztext []string

			if existsArtikel {
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('p.product__manufactorId')).map(p => p.textContent)`, &bestellnummer).Do(ctx)
				if err != nil {
					return err
				}
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('a.product__title')).map(a => a.textContent)`, &kurztext).Do(ctx)
				if err != nil {
					return err
				}
				for a := range bestellnummer {
					product.Bestellnummer = bestellnummer[a]
					product.Kurztext = kurztext[a]
					produkte[bestellnummer[a]] = product
					fmt.Printf("Bestellnummer: %-50s", bestellnummer[a])
					fmt.Printf("\n")
				}

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

func importJSON() {

	jsonFile_LISTE, err := os.Open("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Blame_Rittal1.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()

	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)

	artikelRaw := []Artikel{}
	artikelSeed := map[string]*ArtikelSeed{}
	json.Unmarshal(byteValue_LISTE, &artikelRaw)

	for _, b := range artikelRaw {

		bestellnummer := strings.Split(b.URL, "=")[1]
		artikelnummer := b.Artikelnummer

		setSeed(artikelSeed, "Bestellnummer", bestellnummer, bestellnummer, b.Artikelnummer, b.URL)
		setSeed(artikelSeed, "Artikelnummer", artikelnummer, bestellnummer, b.Artikelnummer, b.URL)
	}

	data, err := json.MarshalIndent(artikelSeed, "", "\t")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Blame_Artikel_Seed_Rittal.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func setSeed(artikelSeed map[string]*ArtikelSeed, prefix string, key string, bestellnummer string, artikelnummer string, url string) {
	key_Clean := prefix + ": " + cleanStringComplete(key)
	bestellnummer_Clean := strings.ToUpper(bestellnummer)
	artikelnummer_Clean := strings.ToUpper(artikelnummer)

	_, okBestellnummer := artikelSeed[key_Clean]
	if !okBestellnummer {
		artikelSeed[key_Clean] = &ArtikelSeed{
			Bestellnummer: bestellnummer_Clean,
			Artikelnummer: artikelnummer_Clean,
			URL:           url,
		}
	} else {
		if len(url) < len(artikelSeed[key_Clean].URL) {
			artikelSeed[key_Clean].URL = url
			fmt.Printf("Correcting " + key)
			fmt.Printf("\n")
			fmt.Printf("Old Artikelnummmer: %-73s", artikelSeed[key_Clean].Artikelnummer)
			fmt.Printf("Old URL: %-50s", artikelSeed[key_Clean].URL)
			fmt.Printf("\n")
			fmt.Printf("New Artikelnummmer: %-73s", artikelnummer_Clean)
			fmt.Printf("New URL: %-50s", url)
			fmt.Printf("\n")
			fmt.Println("--")
		} else {
			fmt.Printf("duplicate of " + key + " | Ignoring ")
			fmt.Printf("\n")
			fmt.Printf("Old Artikelnummmer: %-30s", artikelSeed[key_Clean].Artikelnummer)
			fmt.Printf("Old URL len: %-30d", len(artikelSeed[key_Clean].URL))
			fmt.Printf("Old URL: %-50s", artikelSeed[key_Clean].URL)
			fmt.Printf("\n")
			fmt.Printf("New Artikelnummmer: %-30s", artikelnummer_Clean)
			fmt.Printf("New URL len: %-30d", len(url))
			fmt.Printf("New URL: %-50s", url)
			fmt.Printf("\n")
			fmt.Println("--")
		}

	}
}
func cleanStringComplete(x string) string {
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "/", "")
	x = strings.ReplaceAll(x, ",", "")
	x = strings.ReplaceAll(x, "ü", "ue")
	x = strings.ReplaceAll(x, "ä", "ae")
	x = strings.ReplaceAll(x, "ö", "oe")
	x = strings.ReplaceAll(x, "Ü", "Ue")
	x = strings.ReplaceAll(x, "Ä", "Ae")
	x = strings.ReplaceAll(x, "Ö", "Oe")
	return x
}
