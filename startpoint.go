package main

import (
	"fmt"
	"strings"
)

var Stueckliste1 = []*Artikel{}
var Stueckliste2 = []*Artikel{}

func (a *App) ExcelChoice(file1 string, file2 string) {
	Stueckliste1 = []*Artikel{}
	Stueckliste2 = []*Artikel{}
	fmt.Printf("file1: %-20s", file1)
	fmt.Printf("file2: %-20s", file2)
	fmt.Printf("\n")
	STD_Read_Lagerbestand(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\8000772_Stückliste.xlsx", &Stueckliste1, "KNT")
	STD_Read_Lagerbestand(stueckliste_Kroenert, file2, &Stueckliste2, "KNT")

	for _, value := range Stueckliste1 {
		if strings.Contains(value.Hersteller, "RIT") {
			fmt.Printf("ERP: %-20s", value.ERP)
			fmt.Printf("Stückzahl: %-10.2f", value.Stueckzahl)
			fmt.Printf("Bestellnummer: %-20s", value.Bestellnummer)
			fmt.Printf("Hersteller: %-20s", value.Hersteller)
			fmt.Printf("Bestellnr_L1: %-20s", value.Bestellnr_L1)
			fmt.Printf("Herstellertyp: %-20s", value.Herstellertyp)
			fmt.Printf("Beistellung: %-20s", value.Beistellung)
			fmt.Printf("\n")
		}
	}
	fmt.Println("Start Point Fertig")
}
