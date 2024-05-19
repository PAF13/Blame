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
		artikelnummer := b.Artikelnummer[0]

		setSeed(artikelSeed, "Bestellnummer", bestellnummer, bestellnummer, b.Artikelnummer[0], b.URL)
		setSeed(artikelSeed, "Artikelnummer", artikelnummer, bestellnummer, b.Artikelnummer[0], b.URL)
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

func main() {
	//importJSON()
	//Scraper()
	scraperSiemens()
	fmt.Println("Done")
}

func scraperSiemens() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 99999*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		log.Fatalf("Failed to enable network: %v", err)
	}

	url := "https://mall.industry.siemens.com/mall/de/de/Catalog/Product/3LD2418-0TK13"

	visited := make(map[string]bool)

	scrapeRittal0(url, &ctx, visited)
	//writeJsonFileScrapper("Rittal", artikel)
}
