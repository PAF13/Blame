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
		fmt.Printf("ERP: %-20s", value.ERP)
		fmt.Printf("Stückzahl: %-10.2f", value.Stueckzahl)
		fmt.Printf("Bestellnummer: %-25s", value.Bestellnummer)
		fmt.Printf("Hersteller: %-30s", value.Hersteller)
		fmt.Printf("Bestellnr_L1: %-20s", value.Bestellnr_L1)
		fmt.Printf("Herstellertyp: %-20s", value.Herstellertyp)
		fmt.Printf("Beistellung: %-20s", value.Beistellung)
		fmt.Printf("\n")
	}
}
