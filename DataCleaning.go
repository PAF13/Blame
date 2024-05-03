package main

import (
	"fmt"
	"strings"
)

func STD_Clean_Lagerbestand(lagerbestand []*Artikel) {
	for num, artikel := range lagerbestand {
		if artikel.Bestellnummer != "" {
			vorhanden := false
			artikelnum, ok := lagerbestand_Siteca_Map[artikel.Bestellnummer]
			erpNum := ""
			if !ok {
				for _, ArtikelStamm := range lagerbestand_Siteca {
					if strings.Contains(strings.ToLower(ArtikelStamm.Bestellnummer), strings.ToLower(artikel.Bestellnummer)) {
						erpNum = ArtikelStamm.ERP
						vorhanden = true
						break
					}
				}
			} else {
				erpNum = artikelnum.ERP
				vorhanden = true
			}

			if vorhanden {
				lagerbestand[num].STD_Clean_Lagerbestand_Update(erpNum)
			}
		}
	}
}
func (lagerbestand *Artikel) STD_Clean_Lagerbestand_Update(erp string) {
	lagerbestand.ERP = erp
	lagerbestand.Quelle = "Siteca"
	//lagerbestand.Hersteller = lagerbestand_Siteca_Map[lagerbestand.Bestellnummer].Hersteller
}

func (lagerbestand *Artikel) STD_Stueckliste_Update(stueckzahl float64) {
	lagerbestand.Stueckzahl = lagerbestand.Stueckzahl + stueckzahl
}

func (lagerbestand *Artikel) STD_Stueckliste_Update2(stueckzahl1 float64, stueckzahl2 float64) {
	lagerbestand.Stueckzahl = stueckzahl2 - stueckzahl1
}

func (lagerbestand *Artikel) STD_Stueckliste_Update3(stueckzahl float64) {
	lagerbestand.Stueckzahl = stueckzahl * -1
}

func Sitecavergleich(artikelstammdaten map[string]*Artikel, liste map[string]*Artikel) {
	for artikel, row := range liste {
		stammnum, ok := artikelstammdaten[artikel]
		var artikelstammnum *Artikel
		vorhanden := false
		ERP := ""
		fmt.Println("checking artikel")
		if !ok {
			fmt.Println("artikel not found")
			for artikel2, ArtikelStamm := range artikelstammdaten {
				if ArtikelStamm.Herstellertyp == artikel2 {
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				} else if ArtikelStamm.HerstellerEplan == artikel2 {
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				} else if ArtikelStamm.Bestellnr_L1 == artikel2 {
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				}
			}

		} else {
			fmt.Println("artikel found")
			artikelstammnum = stammnum
			vorhanden = true
		}
		fmt.Println("adding " + artikel + " to list")
		if vorhanden {
			liste[artikel] = &Artikel{
				ERP:                ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: artikelstammnum.ArtikelnummerEplan,
				Hersteller:         artikelstammnum.Hersteller,
				Beschreibung:       artikelstammnum.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             "Siteca",
				Stand:              "Topix verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		} else {
			liste[artikel] = &Artikel{
				ERP:                row.ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: row.ArtikelnummerEplan,
				Hersteller:         row.Hersteller,
				Beschreibung:       row.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             row.Quelle,
				Stand:              "Topix nicht verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		}
	}
}

func Sitecavergleich22(artikelstammdaten map[string]*Artikel, liste map[string]*Artikel, stueckliste *[]*Artikel) {
	for num, row := range *stueckliste {
		//erpNum := ""
		_, ok := artikelstammdaten[row.Bestellnummer]
		if ok {
			(*stueckliste)[num] = &Artikel{
				ERP:                artikelstammdaten[row.Bestellnummer].ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: artikelstammdaten[row.Bestellnummer].ArtikelnummerEplan,
				Hersteller:         artikelstammdaten[row.Bestellnummer].Hersteller,
				Beschreibung:       artikelstammdaten[row.Bestellnummer].Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             "Siteca",
				Stand:              "Topix verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		} else {
			(*stueckliste)[num] = &Artikel{
				ERP:                row.ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: row.ArtikelnummerEplan,
				Hersteller:         row.Hersteller,
				Beschreibung:       row.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             row.Quelle,
				Stand:              "Topix nicht verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		}
	}
}

func Sitecavergleich2(artikelstammdaten map[string]*Artikel, liste map[string]*Artikel, stueckliste *[]*Artikel) {
	for artikel, row := range *stueckliste {
		stammnum, ok := artikelstammdaten[row.Bestellnummer]
		var artikelstammnum *Artikel
		vorhanden := false
		ERP := ""
		fmt.Println("checking artikel")
		if !ok {
			fmt.Println("artikel not found")
			for artikel2, ArtikelStamm := range artikelstammdaten {
				if ArtikelStamm.Herstellertyp == artikel2 {
					fmt.Println("artikel found in Herstellertyp")
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				} else if ArtikelStamm.HerstellerEplan == artikel2 {
					fmt.Println("artikel found in HerstellerEplan")
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				} else if ArtikelStamm.Bestellnr_L1 == artikel2 {
					fmt.Println("artikel found in Bestellnr_L1")
					vorhanden = true
					ERP = ArtikelStamm.ERP
					artikelstammnum = ArtikelStamm
				}
			}

		} else {
			fmt.Println("artikel found")
			artikelstammnum = stammnum
			vorhanden = true
			ERP = artikelstammdaten[row.Bestellnummer].ERP
		}
		fmt.Println("adding " + (*stueckliste)[artikel].Bestellnummer + " to list")
		if vorhanden {
			(*stueckliste)[artikel] = &Artikel{
				ERP:                ERP,
				Bestellnummer:      artikelstammnum.Bestellnummer,
				ArtikelnummerEplan: artikelstammnum.ArtikelnummerEplan,
				Hersteller:         artikelstammnum.Hersteller,
				Beschreibung:       artikelstammnum.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             "Siteca",
				Stand:              "Topix verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		} else {
			(*stueckliste)[artikel] = &Artikel{
				ERP:                "fehlt",
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: row.ArtikelnummerEplan,
				Hersteller:         row.Hersteller,
				Beschreibung:       row.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             row.Quelle,
				Stand:              "Topix nicht verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		}
	}
}

func STD_Sum(list *[]*Artikel, listsum *[]*Artikel, listmap map[string]*Artikel) {
	for _, b := range *list {
		_, ok := listmap[b.Ort+b.Bestellnummer]

		if ok {
			listmap[b.Ort+b.Bestellnummer].STD_Stueckliste_Update(b.Stueckzahl)
		} else {
			listmap[b.Ort+b.Bestellnummer] = b
		}
	}

	for _, b := range listmap {
		*listsum = append(*listsum, b)
	}
}
