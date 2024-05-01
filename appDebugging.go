package main

import (
	"fmt"
)

func (a *App) Lagerbestand(lager string) {
	fmt.Println("Lagerbestand")
	if lager == "kreonert" {

	}
	if lager == "siteca" {

	}
	fmt.Println("Fertig")
}

func printLagerbestand(lager map[string]*Artikel) {
	for _, value := range lager {
		fmt.Printf("ERP: %-40s", value.ERP)
		fmt.Printf("Artikelnummer: %-50s", value.ArtikelnummerEplan)
		fmt.Printf("Bestellnummer: %-40s", value.Bestellnummer)
		fmt.Printf("Hersteller: %-40s", value.Hersteller)
		fmt.Printf("Stückzahl: %-30f", value.Stueckzahl)
		fmt.Printf("Typ: %-40s", value.Typ)
		fmt.Printf("\n")
	}
}
