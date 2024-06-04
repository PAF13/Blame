package parts

import "strconv"

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
		keyBetriebsmittel := "==" + safeHeader(rows[i], headersClean["FunktionaleZuordnung"]) +
			"=" + safeHeader(rows[i], headersClean["Funktionskennzeichen"]) +
			"++" + safeHeader(rows[i], headersClean["Aufstellungsort"]) +
			"+" + bestellnummerCleaner(safeHeader(rows[i], headersClean["Ortskennzeichen"])) +
			"-" + bestellnummerCleaner(safeHeader(rows[i], headersClean["BMK"])) +
			"|" + bestellnummerCleaner(safeHeader(rows[i], headersClean["Bestellnummer"]))
		keyFilter := "==" + safeHeader(rows[i], headersClean["FunktionaleZuordnung"]) +
			"+" + bestellnummerCleaner(safeHeader(rows[i], headersClean["Ortskennzeichen"]))
		//keyBetriebsmittel := "==" + safeHeader(rows[i], headersClean["FunktionaleZuordnung"]) + "+" + safeHeader(rows[i], headersClean["Ortskennzeichen"]) + "|" + bestellnummerCleaner(safeHeader(rows[i], headersClean["Ortskennzeichen"]))
		if keyBetriebsmittel != "==+|" {
			_, okFilter := filter.Filter[keyFilter]
			if !okFilter {
				if safeHeader(rows[i], headersClean["Beistellung"]) == "Siteca" {
					filter.Filter[keyFilter] = true
				} else {
					filter.Filter[keyFilter] = false
				}

			} else {
				if safeHeader(rows[i], headersClean["Beistellung"]) == "Siteca" {
					filter.Filter[keyFilter] = true
				}
			}
			_, okBetriebsmittel := Liste.Betriebsmittel[keyBetriebsmittel]
			if okBetriebsmittel {
				Liste.Betriebsmittel[keyBetriebsmittel].Artikel = append(Liste.Betriebsmittel[keyBetriebsmittel].Artikel, NewArtikelTemp(headersClean, rows[i]))
			} else {
				Liste.Betriebsmittel[keyBetriebsmittel] = NewBetriebsmittelTemp(headersClean, rows[i])
				Liste.Betriebsmittel[keyBetriebsmittel].Artikel = append(Liste.Betriebsmittel[keyBetriebsmittel].Artikel, NewArtikelTemp(headersClean, rows[i]))
			}
		}
	}
}
