package parts

import "strconv"

func NewBetriebsmittelTemp(headersClean map[string]uint64, row []string) *BETRIEBSMITELL {
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

func NewArtikelTemp(headersClean map[string]uint64, row []string) *ARTIKEL {
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

func NewBetriebsmittel() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{}
}
func NewArtikel() *ARTIKEL {
	return &ARTIKEL{}
}

func NewLagerliste() *LAGERLISTE {
	return &LAGERLISTE{}
}

func NewBetriebsmillliste() *BETRIEBSMITELLLISTE {
	return &BETRIEBSMITELLLISTE{
		Betriebsmittel: make(map[string]*BETRIEBSMITELL),
	}
}
