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
	bestellnummer := []string{
		"3LD2418-0TK13",
	}
	switch hersteller {
	case "SIEMENS":
		scraperSiemens(bestellnummer)
	}

	fmt.Println("Done")
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

func scraperSiemens(bestellnummer []string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 99999*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		log.Fatalf("Failed to enable network: %v", err)
	}
	var produkte map[string]*Product = make(map[string]*Product)
	for _, b := range bestellnummer {
		scrapesSiemensLoop(b, &ctx, produkte)
	}

	writeJsonFileScrapper("Rittal", produkte)
}
func scrapesSiemensLoop(bestellnummer string, ctx *context.Context, produkte map[string]*Product) {
	url := "https://mall.industry.siemens.com/mall/de/de/Catalog/Product/" + bestellnummer
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

			err := chromedp.EvaluateAsDevTools(`document.querySelector('span.productIdentifier') !== null`, &existsArtikel).Do(ctx)
			if err != nil {
				return err
			}
			var bestellnummer []string
			var artikelnummer []string
			var verrpackungsmasse string
			var werte []*[]string = []*[]string{
				&bestellnummer,
				&artikelnummer,
			}

			if existsArtikel {
				log.Println("On Artikel page")
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('div.productDataGroup span.productIdentifier')).map(span => span.textContent)`, &bestellnummer).Do(ctx)
				err = chromedp.Evaluate(`Array.from(document.querySelectorAll('div.productDataGroup span.productIdentifier')).map(span => span.textContent)`, &artikelnummer).Do(ctx)
				err = chromedp.Evaluate(`(function() {
					let label = Array.from(document.querySelectorAll('td.productDetailsTable_DataLabel'))
						.find(el => el.textContent.trim() === 'Verpackungsmaße');
					return label ? label.nextElementSibling.textContent.trim() : '';
				})()`, &verrpackungsmasse).Do(ctx)
				if err != nil {
					return err
				}
				for _, b := range werte {
					if len(*b) > 1 {
						log.Println("Error: more than 1 element found | ")
					}
				}
				product.Bestellnummer = bestellnummer[0]
				product.Artikelnummer = artikelnummer[0]
				product.Produktinformation.Material = verrpackungsmasse
				produkte[product.Bestellnummer] = product
				fmt.Printf("Bestellnummer: %-20s", product.Bestellnummer)
				fmt.Printf("URL: %-50s", product.URL)
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
