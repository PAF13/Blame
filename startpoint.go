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
	StuecklisteDif := []*Artikel{}

	fmt.Printf("file1: %-20s", file1)
	fmt.Printf("file2: %-20s", file2)
	fmt.Printf("\n")
	STD_Read_Lagerbestand(stueckliste_projekt, file1, &Stueckliste1, "KNT")
	STD_Clean_Lagerbestand(Stueckliste1)
	STD_Sum(&Stueckliste1, &Stueckliste1_Sum, Stueckliste1_Map)

	STD_Read_Lagerbestand(stueckliste_projekt, file2, &Stueckliste2, "KNT")
	STD_Clean_Lagerbestand(Stueckliste2)
	STD_Sum(&Stueckliste2, &Stueckliste2_Sum, Stueckliste2_Map)

	for artikelnum, artikel := range Stueckliste1_Map {
		_, ok := Stueckliste2_Map[artikelnum]
		if ok {
			if Stueckliste1_Map[artikelnum].Stueckzahl == Stueckliste2_Map[artikelnum].Stueckzahl {
				delete(Stueckliste1_Map, artikelnum)
			} else {
				artikel.STD_Stueckliste_Update2(Stueckliste1_Map[artikelnum].Stueckzahl, Stueckliste2_Map[artikelnum].Stueckzahl)
			}
			delete(Stueckliste2_Map, artikelnum)
		}
	}
	for artikelnum, artikel := range Stueckliste2_Map {
		Stueckliste1_Map[artikelnum] = artikel
	}

	fmt.Println("Stueckliste diff")
	for _, artikel := range Stueckliste1_Map {
		StuecklisteDif = append(StuecklisteDif, artikel)
	}

	STD_Write_Stueckliste("Stueckliste_Dif.xlsx", StuecklisteDif)

	fmt.Println("Start Point Fertig")
}

func (a *App) StuecklisteSum(file string) {

	Stueckliste := []*Artikel{}
	Stueckliste2 := []*Artikel{}
	Stueckliste_Map := make(map[string]*Artikel)

	STD_Read_Lagerbestand(stueckliste_projekt, file, &Stueckliste, "")
	STD_Write_Stueckliste("Liste_Stueckliste.xlsx", Stueckliste)
	STD_Clean_Lagerbestand(Stueckliste)
	STD_Write_Stueckliste("Liste_Stueckliste_Clean.xlsx", Stueckliste)

	STD_Sum(&Stueckliste, &Stueckliste2, Stueckliste_Map)

	STD_Write_Stueckliste("Stueckliste_Sum.xlsx", Stueckliste2)

	fmt.Println("StuecklisteSum finished")
}
