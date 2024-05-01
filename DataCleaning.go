package main

import (
	"fmt"
	"strings"
)

func STD_Clean_Lagerbestand(lagerbestand []*Artikel) {
	for num, artikel := range lagerbestand {
		vorhanden := false

		erp := artikel.ERP
		_, ok := lagerbestand_Siteca_Map[artikel.Bestellnummer]
		for _, ArtikelStamm1 := range lagerbestand_Siteca_Map {
			if ok {

				erp = lagerbestand_Siteca_Map[artikel.Bestellnummer].ERP
				vorhanden = true
			} else if strings.Contains(ArtikelStamm1.Bestellnr_L1, artikel.Bestellnummer) {
				erp = ArtikelStamm1.ERP
				vorhanden = true
			} else if strings.Contains(ArtikelStamm1.Herstellertyp, artikel.Bestellnummer) {
				erp = ArtikelStamm1.ERP
				vorhanden = true
			}
		}
		if vorhanden {
			lagerbestand[num].STD_Clean_Lagerbestand_Update(erp)
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
