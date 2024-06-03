package parts

import (
	"math"
	"strconv"
)

func NewBetriebsmittelTemp(headersClean map[string]uint64, row []string) *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: safeHeader(row, headersClean["FunktionaleZuordnung"]),
			Funktionskennzeichen: safeHeader(row, headersClean["Funktionskennzeichen"]),
			Aufstellungsort:      safeHeader(row, headersClean["Aufstellungsort"]),
			Ortskennzeichen:      safeHeader(row, headersClean["Ortskennzeichen"]),
			BMK:                  safeHeader(row, headersClean["BMK"]),
		},
		Artikel: []*ARTIKEL{},
	}
}
func NewBetriebsmittelTemp2(headersClean map[string]uint64, row []string) *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: safeHeader(row, headersClean["FunktionaleZuordnung"]),
			Ortskennzeichen:      safeHeader(row, headersClean["Ortskennzeichen"]),
			/*FunktionaleZuordnung: safeHeader(row, headersClean["FunktionaleZuordnung"]),
			Funktionskennzeichen: safeHeader(row, headersClean["Funktionskennzeichen"]),
			Aufstellungsort:      safeHeader(row, headersClean["Aufstellungsort"]),
			Ortskennzeichen:      safeHeader(row, headersClean["Ortskennzeichen"]),
			BMK:                  safeHeader(row, headersClean["BMK"]),*/
		},
		Artikel: []*ARTIKEL{},
	}
}
func NewArtikelTemp(headersClean map[string]uint64, row []string) *ARTIKEL {
	stueckzahl, _ := strconv.ParseFloat(safeHeader(row, headersClean["Stueckzahl"]), 64)
	stueckzahlMoeller, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Moeller"]), 64)
	stueckzahlKNT, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_KNT"]), 64)
	stueckzahlSiteca, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Siteca"]), 64)
	return &ARTIKEL{
		Bestellnummer:      safeHeader(row, headersClean["Bestellnummer"]),
		ERP:                safeHeader(row, headersClean["ERP"]),
		ERP_KNT:            safeHeader(row, headersClean["ERP_KNT"]),
		Hersteller:         safeHeader(row, headersClean["Hersteller"]),
		Stueckzahl:         stueckzahl,
		Bestellung_Moeller: stueckzahlMoeller,
		Bestellung_KNT:     stueckzahlKNT,
		Bestellung_Siteca:  stueckzahlSiteca,
		Beschreibung:       safeHeader(row, headersClean["Beschreibung"]),
		Beistellung:        safeHeader(row, headersClean["Beistellung"]),
		Funktionsgruppe:    safeHeader(row, headersClean["Funktionsgruppe"]),
	}
}
func safeHeader(row []string, column uint64) string {
	var value string
	if column == math.MaxUint64 {
		value = ""
	} else {
		value = row[column]
	}
	return value
}
func NewBetriebsmittel() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{}
}
func NewArtikel() *ARTIKEL {
	return &ARTIKEL{}
}

func NewLagerliste() *LAGERLISTE {
	return &LAGERLISTE{
		Betriebsmittel: make(map[string]*ARTIKEL),
	}
}

func NewBetriebsmillliste() *BETRIEBSMITELLLISTE {
	return &BETRIEBSMITELLLISTE{
		Betriebsmittel: make(map[string]*BETRIEBSMITELL),
	}
}
