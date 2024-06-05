package parts

import (
	"fmt"
	"strconv"
	"strings"
)

func (Liste *LAGERLISTE) setListe(P_rows *[][]string, headersClean map[string]uint64, headerRow int) {
	rows := *P_rows
	for i := headerRow + 1; i < len(rows)-1; i++ {
		if bestellnummerCleaner(safeHeader(rows[i], headersClean["Bestellnummer"])) != "" {
			keyBetriebsmittel := bestellnummerCleaner(safeHeader(rows[i], headersClean["Bestellnummer"]))

			_, okBetriebsmittel := Liste.Betriebsmittel[keyBetriebsmittel]
			if okBetriebsmittel {
				stueckzahlBeistellung, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Stueckzahl"]), 64)
				stueckzahlMoeller, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Bestellung_Moeller"]), 64)
				stueckzahlSiteca, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Lager_Siteca"]), 64)
				stueckzahlKNT, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Bestellung_KNT"]), 64)
				if Liste.Betriebsmittel[keyBetriebsmittel].Hersteller == "" {
					Liste.Betriebsmittel[keyBetriebsmittel].Hersteller = safeHeader(rows[i], headersClean["Hersteller"])
				}
				if Liste.Betriebsmittel[keyBetriebsmittel].ERP == "" {
					Liste.Betriebsmittel[keyBetriebsmittel].ERP = safeHeader(rows[i], headersClean["ERP"])
				}
				if Liste.Betriebsmittel[keyBetriebsmittel].ERP_KNT == "" {
					Liste.Betriebsmittel[keyBetriebsmittel].ERP_KNT = safeHeader(rows[i], headersClean["ERP_KNT"])
				}
				Liste.Betriebsmittel[keyBetriebsmittel].Bestellung_Moeller = Liste.Betriebsmittel[keyBetriebsmittel].Bestellung_Moeller + stueckzahlMoeller
				Liste.Betriebsmittel[keyBetriebsmittel].Bestellung_KNT = Liste.Betriebsmittel[keyBetriebsmittel].Bestellung_KNT + stueckzahlKNT
				Liste.Betriebsmittel[keyBetriebsmittel].Lager_Siteca = Liste.Betriebsmittel[keyBetriebsmittel].Lager_Siteca + stueckzahlSiteca
				if safeHeader(rows[i], headersClean["Funktionsgruppe"]) == "Beistellung" {
					Liste.Betriebsmittel[keyBetriebsmittel].Beistellung_Stueckzahl = Liste.Betriebsmittel[keyBetriebsmittel].Beistellung_Stueckzahl + stueckzahlBeistellung
				}
				if Liste.Betriebsmittel[keyBetriebsmittel].ERP == "" {
					safeHeader(rows[i], headersClean["ERP"])
				}

			} else {
				Liste.Betriebsmittel[keyBetriebsmittel] = NewArtikelTemp2(headersClean, rows[i])
			}
		}
	}

}

func (Liste *BETRIEBSMITELLLISTE) setListe(P_rows *[][]string, headersClean map[string]uint64, headerRow int, zusatzfilter bool, filter *FILTER) {
	rows := *P_rows
	for i := headerRow + 1; i < len(rows)-1; i++ {
		keyFilter, keyBetriebsmittel, quelle := filter.setProdukte(rows[i], headersClean)
		fmt.Println(keyFilter)
		_, okFilter := filter.Filter[keyBetriebsmittel]
		if !okFilter {
			filter.Filter[keyBetriebsmittel] = true
		}
		if (zusatzfilter && safeHeader(rows[i], headersClean["Funktionsgruppe"]) == "Beistellung") || !zusatzfilter {
			_, okBetriebsmittel := Liste.Betriebsmittel[keyBetriebsmittel]
			if okBetriebsmittel {
				Liste.Betriebsmittel[keyBetriebsmittel].Artikel = append(Liste.Betriebsmittel[keyBetriebsmittel].Artikel, NewArtikelTemp(headersClean, rows[i], quelle))
			} else {
				Liste.Betriebsmittel[keyBetriebsmittel] = NewBetriebsmittelTemp(headersClean, rows[i])
				Liste.Betriebsmittel[keyBetriebsmittel].Artikel = append(Liste.Betriebsmittel[keyBetriebsmittel].Artikel, NewArtikelTemp(headersClean, rows[i], quelle))
			}
		}

	}
}

func (filter *FILTER) setProdukte(row []string, headersClean map[string]uint64) (string, string, string) {
	FunktionaleZuordnung := "==" + safeHeader(row, headersClean["FunktionaleZuordnung"])
	//Funktionskennzeichen := "=" + safeHeader(row, headersClean["Funktionskennzeichen"])
	Aufstellungsort := "++" + safeHeader(row, headersClean["Aufstellungsort"])
	Ortskennzeichen := "+" + bestellnummerCleaner(safeHeader(row, headersClean["Ortskennzeichen"]))
	BMK := "-" + bestellnummerCleaner(safeHeader(row, headersClean["BMK"]))
	//Funktionsgruppe := safeHeader(row, headersClean["Funktionsgruppe"])

	var ERP string
	var keyProduct string
	var keyStueck string
	var producttype string
	if bestellnummerCleaner(safeHeader(row, headersClean["Bestellnummer"])) == "" {
		ERP = "|" + safeHeader(row, headersClean["ERP_KNT"])
	} else {
		ERP = "|" + bestellnummerCleaner(safeHeader(row, headersClean["Bestellnummer"]))
	}
	switch {
	case strings.Contains(Ortskennzeichen, "EC") || strings.Contains(Ortskennzeichen, "EB"):
		keyProduct = FunktionaleZuordnung + Ortskennzeichen
		keyStueck = FunktionaleZuordnung + Ortskennzeichen + ERP
		filter.Product[keyProduct] = "Schaltschrank|" + keyProduct
		producttype = filter.Product[keyProduct]
	case strings.Contains(Ortskennzeichen, "EP"):
		keyProduct = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen
		keyStueck = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen + ERP
		filter.Product[keyProduct] = "Bedienpult|" + keyProduct
		producttype = filter.Product[keyProduct]
	case repschalt(Ortskennzeichen):
		keyProduct = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen + BMK
		keyStueck = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen + BMK + ERP
		filter.Product[keyProduct] = "Reparaturschalter|" + keyProduct
		producttype = filter.Product[keyProduct]
	default:
		keyProduct = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen + BMK
		keyStueck = FunktionaleZuordnung + Aufstellungsort + Ortskennzeichen + BMK + ERP
		filter.Product[keyProduct] = "Missing|" + keyProduct
		producttype = filter.Product[keyProduct]
	}
	return keyProduct, keyStueck, producttype
}
func repschalt(Ortskennzeichen string) bool {
	checker := []bool{
		strings.Contains(Ortskennzeichen, "PU"),
		strings.Contains(Ortskennzeichen, "RW"),
		strings.Contains(Ortskennzeichen, "ST"),
		strings.Contains(Ortskennzeichen, "UW"),
		strings.Contains(Ortskennzeichen, "CS"),
		strings.Contains(Ortskennzeichen, "DS"),
		strings.Contains(Ortskennzeichen, "WI"),
		strings.Contains(Ortskennzeichen, "MH"),
		strings.Contains(Ortskennzeichen, "DR"),
		strings.Contains(Ortskennzeichen, "CH"),
		strings.Contains(Ortskennzeichen, "MF"),
	}
	var ok bool
	for _, b := range checker {
		if b {
			ok = true
			break
		}
	}
	return ok
}
