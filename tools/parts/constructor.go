package parts

import "strconv"

func NewBetriebsmittel(headersClean map[string]uint64, row []string) *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: row[headersClean["FunktionaleZuordnung"]],
			Funktionskennzeichen: row[headersClean["Funktionskennzeichen"]],
			Aufstellungsort:      row[headersClean["Aufstellungsort"]],
			Ortskennzeichen:      row[headersClean["Ortskennzeichen"]],
			BMK:                  row[headersClean["BMK"]],
		},
		Artikel: []*ARTIKEL{},
	}
}

func NewArtikel(headersClean map[string]uint64, row []string) *ARTIKEL {
	stueckzahl, _ := strconv.ParseFloat(row[headersClean["Stueckzahl"]], 64)
	return &ARTIKEL{
		Bestellnummer: bestellnummerCleaner(row[headersClean["Bestellnummer"]]),
		ERP:           "",
		ERP_KNT:       row[headersClean["ERP_KNT"]],
		Hersteller:    row[headersClean["Hersteller"]],
		Stueckzahl:    stueckzahl,
		Beschreibung:  row[headersClean["Beschreibung"]],
	}
}
