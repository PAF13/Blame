package parts

import (
	"strconv"
	"strings"
)

func (Liste *LAGERLISTE) setListe(P_rows *[][]string, headersClean map[string]uint64, headerRow int, fileName string) {
	rows := *P_rows
	for i := headerRow + 1; i < len(rows)-1; i++ {
		if bestellnummerCleaner(safeHeader(rows[i], headersClean["Bestellnummer"])) != "" {
			keyBetriebsmittel := bestellnummerCleaner(safeHeader(rows[i], headersClean["Bestellnummer"]))

			_, okBetriebsmittel := Liste.Artikel[keyBetriebsmittel]
			if !okBetriebsmittel {
				Liste.Artikel[keyBetriebsmittel] = NewArtikelTemp(headersClean, rows[i], "")
			} else {
				if Liste.Artikel[keyBetriebsmittel].ERP == "" {
					Liste.Artikel[keyBetriebsmittel].ERP = safeHeader(rows[i], headersClean["ERP"])
				}
				if Liste.Artikel[keyBetriebsmittel].ERP_KNT == "" {
					Liste.Artikel[keyBetriebsmittel].ERP_KNT = safeHeader(rows[i], headersClean["ERP_KNT"])
				}
				stueckzahlMoeller, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Lager_Moeller"]), 64)
				stueckzahlKNT, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Lager_KNT"]), 64)
				stueckzahlSiteca, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["Lager_Siteca"]), 64)
				EKKNT, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["EK_KNT"]), 64)
				EKSiteca, _ := strconv.ParseFloat(safeHeader(rows[i], headersClean["EK_Siteca"]), 64)
				Liste.Artikel[keyBetriebsmittel].Bestellung_Moeller = Liste.Artikel[keyBetriebsmittel].Bestellung_Moeller + stueckzahlMoeller
				Liste.Artikel[keyBetriebsmittel].Bestellung_KNT = Liste.Artikel[keyBetriebsmittel].Bestellung_KNT + stueckzahlKNT
				Liste.Artikel[keyBetriebsmittel].Lager_Siteca = Liste.Artikel[keyBetriebsmittel].Bestellung_Siteca + stueckzahlSiteca
				if Liste.Artikel[keyBetriebsmittel].EK_KNT < EKKNT {
					Liste.Artikel[keyBetriebsmittel].EK_KNT = EKKNT
				}
				if Liste.Artikel[keyBetriebsmittel].EK_Siteca < EKSiteca {
					Liste.Artikel[keyBetriebsmittel].EK_Siteca = EKSiteca
				}
			}

			switch fileName {
			case "TOPIX":
				Liste.Artikel[keyBetriebsmittel].DataSource.Siteca = true
			case "MOELLER":
				Liste.Artikel[keyBetriebsmittel].DataSource.Moeller = true
			case "LAGERHUETER":
				Liste.Artikel[keyBetriebsmittel].DataSource.KNT = true
			}
		}
	}
}

func (Liste *LAGERLISTE) setListe2(part []*ARTIKEL) {

	for _, b := range part {
		keyBetriebsmittel := bestellnummerCleaner(b.Bestellnummer)

		_, okBetriebsmittel := Liste.Artikel[keyBetriebsmittel]
		if !okBetriebsmittel {
			Liste.Artikel[keyBetriebsmittel] = b
		} else {
			if Liste.Artikel[keyBetriebsmittel].ERP == "" {
				Liste.Artikel[keyBetriebsmittel].ERP = b.ERP
			}

			if Liste.Artikel[keyBetriebsmittel].ArtikelnummerEplan == "" {
				Liste.Artikel[keyBetriebsmittel].ArtikelnummerEplan = b.ArtikelnummerEplan
			}
			Liste.Artikel[keyBetriebsmittel].ARTIKELINFO = b.ARTIKELINFO
			Liste.Artikel[keyBetriebsmittel].DataSource.Eplan = true
		}
	}

}
func (Liste *BETRIEBSMITELLLISTE) setListe(P_rows *[][]string, headersClean map[string]uint64, headerRow int) {
	rows := *P_rows
	Liste.Produkte = make(map[string]*BETRIEBSMITELLKENNZEICHEN)

	for i := headerRow + 1; i < len(rows); i++ {
		keyBetriebsmittel, quelle := Liste.setProdukte(rows[i], headersClean)
		//_, keyBetriebsmittel, quelle := filter.setProdukte(rows[i], headersClean)
		_, okBetriebsmittel := Liste.Betriebsmittel[keyBetriebsmittel]
		if !okBetriebsmittel {
			Liste.Betriebsmittel[keyBetriebsmittel] = NewBetriebsmittelTemp(headersClean, rows[i])
		}
		Liste.Betriebsmittel[keyBetriebsmittel].Artikel = append(Liste.Betriebsmittel[keyBetriebsmittel].Artikel, NewArtikelTemp(headersClean, rows[i], quelle))
	}
}

func (Liste *BETRIEBSMITELLLISTE) setProdukte(row []string, headersClean map[string]uint64) (string, string) {
	FunktionaleZuordnung := "==" + safeHeader(row, headersClean["FunktionaleZuordnung"])
	Funktionskennzeichen := "=" + safeHeader(row, headersClean["Funktionskennzeichen"])
	Aufstellungsort := "++" + safeHeader(row, headersClean["Aufstellungsort"])
	Ortskennzeichen := "+" + safeHeader(row, headersClean["Ortskennzeichen"])
	BMK := "-" + safeHeader(row, headersClean["BMK"])
	BMKVoll := FunktionaleZuordnung + Funktionskennzeichen + Aufstellungsort + Ortskennzeichen + BMK

	var ERP string
	if bestellnummerCleaner(safeHeader(row, headersClean["Bestellnummer"])) == "" {
		ERP = "|" + safeHeader(row, headersClean["ERP_KNT"])
	} else {
		ERP = "|" + bestellnummerCleaner(safeHeader(row, headersClean["Bestellnummer"]))
	}

	_, ok := Liste.Produkte[BMKVoll]
	if !ok {
		Liste.Produkte[BMKVoll] = &BETRIEBSMITELLKENNZEICHEN{}

		if Liste.ProduktDef.FunktionaleZuordnung {
			Liste.Produkte[BMKVoll].FunktionaleZuordnung = FunktionaleZuordnung
			Liste.Produkte[BMKVoll].BMKVollständig = Liste.Produkte[BMKVoll].BMKVollständig + FunktionaleZuordnung
		}
		if Liste.ProduktDef.Funktionskennzeichen {
			Liste.Produkte[BMKVoll].Funktionskennzeichen = Funktionskennzeichen
			Liste.Produkte[BMKVoll].BMKVollständig = Liste.Produkte[BMKVoll].BMKVollständig + Funktionskennzeichen
		}
		if Liste.ProduktDef.Aufstellungsort {
			Liste.Produkte[BMKVoll].Aufstellungsort = Aufstellungsort
			Liste.Produkte[BMKVoll].BMKVollständig = Liste.Produkte[BMKVoll].BMKVollständig + Aufstellungsort
		}
		if Liste.ProduktDef.Ortskennzeichen {
			Liste.Produkte[BMKVoll].Ortskennzeichen = Ortskennzeichen
			Liste.Produkte[BMKVoll].BMKVollständig = Liste.Produkte[BMKVoll].BMKVollständig + Ortskennzeichen
		}
		if Liste.ProduktDef.Dokumentenart {
			Liste.Produkte[BMKVoll].Dokumentenart = ""
		}
		if Liste.ProduktDef.BenutzerdefinierteStruktur {
			Liste.Produkte[BMKVoll].BenutzerdefinierteStruktur = ""
		}
		if Liste.ProduktDef.Anlagennummer {
			Liste.Produkte[BMKVoll].Anlagennummer = ""
		}
		if Liste.ProduktDef.BMK {
			Liste.Produkte[BMKVoll].BMK = BMK
			Liste.Produkte[BMKVoll].BMKVollständig = Liste.Produkte[BMKVoll].BMKVollständig + BMK
		}

	}
	return Liste.Produkte[BMKVoll].BMKVollständig + ERP, Liste.Produkte[BMKVoll].BMKVollständig
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
