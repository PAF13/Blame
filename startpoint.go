package main

import (
	"fmt"
)

func (a *App) ExcelChoice(file1 string, file2 string) {
	Stueckliste1 := []*Artikel{}
	Stueckliste1_Map := make(map[string]*Artikel)
	Stueckliste1_Sum := []*Artikel{}
	Stueckliste2 := []*Artikel{}
	Stueckliste2_Map := make(map[string]*Artikel)
	Stueckliste2_Sum := []*Artikel{}
	//Stueckliste2_Dif := []*Artikel{}

	fmt.Printf("file1: %-20s", file1)
	fmt.Printf("file2: %-20s", file2)
	fmt.Printf("\n")
	STD_Read_Lagerbestand(stueckliste_projekt, file1, &Stueckliste1, "KNT")
	STD_Clean_Lagerbestand(Stueckliste1)
	STD_Sum(&Stueckliste1, &Stueckliste1_Sum, Stueckliste1_Map)

	STD_Read_Lagerbestand(stueckliste_projekt, file2, &Stueckliste2, "KNT")
	STD_Clean_Lagerbestand(Stueckliste2)
	STD_Sum(&Stueckliste2, &Stueckliste2_Sum, Stueckliste2_Map)

	for _, value := range Stueckliste1_Sum {
		fmt.Printf("ERP: %-20s", value.ERP)
		fmt.Printf("Stückzahl: %-10.2f", value.Stueckzahl)
		fmt.Printf("Bestellnummer: %-25s", value.Bestellnummer)
		fmt.Printf("Hersteller: %-30s", value.Hersteller)
		fmt.Printf("Bestellnr_L1: %-20s", value.Bestellnr_L1)
		fmt.Printf("Herstellertyp: %-20s", value.Herstellertyp)
		fmt.Printf("Beistellung: %-20s", value.Beistellung)
		fmt.Printf("\n")
	}

	fmt.Println("Start Point Fertig")
}
