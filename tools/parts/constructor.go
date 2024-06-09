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
func (liste *BETRIEBSMITELL) NewBetriebsmittelTemp2() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: liste.BMK.FunktionaleZuordnung,
			Funktionskennzeichen: liste.BMK.Funktionskennzeichen,
			Aufstellungsort:      liste.BMK.Aufstellungsort,
			Ortskennzeichen:      liste.BMK.Ortskennzeichen,
			BMK:                  liste.BMK.BMK,
		},
		Artikel: []*ARTIKEL{},
	}
}
func (liste *BETRIEBSMITELL) NewBetriebsmittelTemp3() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: liste.BMK.FunktionaleZuordnung,
			//Funktionskennzeichen: liste.BMK.Funktionskennzeichen,
			//Aufstellungsort:      liste.BMK.Aufstellungsort,
			Ortskennzeichen: liste.BMK.Ortskennzeichen,
			//BMK:                  liste.BMK.BMK,
		},
		Artikel: []*ARTIKEL{},
	}
}
func (liste *BETRIEBSMITELL) NewBetriebsmittelTemp4() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		BMK: BETRIEBSMITELLKENNZEICHEN{
			FunktionaleZuordnung: liste.BMK.FunktionaleZuordnung,
			//Funktionskennzeichen: liste.BMK.Funktionskennzeichen,
			Aufstellungsort: liste.BMK.Aufstellungsort,
			Ortskennzeichen: liste.BMK.Ortskennzeichen,
			//BMK:                  liste.BMK.BMK,
		},
		Artikel: []*ARTIKEL{},
	}
}
func NewArtikelTemp(headersClean map[string]uint64, row []string, quelle string) *ARTIKEL {
	var beistellung bool
	if safeHeader(row, headersClean["Funktionsgruppe"]) == "Beistellung" {
		beistellung = true

	}

	stueckzahl, _ := strconv.ParseFloat(safeHeader(row, headersClean["Stueckzahl"]), 64)
	stueckzahlMoeller, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Moeller"]), 64)
	stueckzahlKNT, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_KNT"]), 64)
	stueckzahlSiteca, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Siteca"]), 64)
	EKKNT, _ := strconv.ParseFloat(safeHeader(row, headersClean["EK_KNT"]), 64)
	EKSiteca, _ := strconv.ParseFloat(safeHeader(row, headersClean["EK_Siteca"]), 64)
	return &ARTIKEL{
		Bestellnummer:      bestellnummerCleaner2(safeHeader(row, headersClean["Bestellnummer"])),
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
		Beigestellt:        beistellung,
		Quelle:             quelle,
		EK_KNT:             EKKNT,
		EK_Siteca:          EKSiteca,
	}
}

func NewArtikelTemp2(headersClean map[string]uint64, row []string) *ARTIKEL {
	var stueckzahl float64
	if safeHeader(row, headersClean["Stueckzahl"]) == "Beistellung" {
		stueckzahl, _ = strconv.ParseFloat(safeHeader(row, headersClean["Stueckzahl"]), 64)
	}

	stueckzahlMoeller, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Moeller"]), 64)
	stueckzahlKNT, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_KNT"]), 64)
	stueckzahlSiteca, _ := strconv.ParseFloat(safeHeader(row, headersClean["Bestellung_Siteca"]), 64)

	return &ARTIKEL{
		Bestellnummer:          safeHeader(row, headersClean["Bestellnummer"]),
		ERP:                    safeHeader(row, headersClean["ERP"]),
		ERP_KNT:                safeHeader(row, headersClean["ERP_KNT"]),
		Hersteller:             safeHeader(row, headersClean["Hersteller"]),
		Beistellung_Stueckzahl: stueckzahl,
		Bestellung_Moeller:     stueckzahlMoeller,
		Bestellung_KNT:         stueckzahlKNT,
		Bestellung_Siteca:      stueckzahlSiteca,
		Beschreibung:           safeHeader(row, headersClean["Beschreibung"]),
		Beistellung:            safeHeader(row, headersClean["Beistellung"]),
		Funktionsgruppe:        safeHeader(row, headersClean["Funktionsgruppe"]),
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
		Artikel: make(map[string]*ARTIKEL),
	}
}

func NewBetriebsmillliste() *BETRIEBSMITELLLISTE {
	return &BETRIEBSMITELLLISTE{
		Betriebsmittel: make(map[string]*BETRIEBSMITELL),
	}
}

func NewFilter() *FILTER {
	return &FILTER{
		Filter:  make(map[string]bool),
		Product: make(map[string]string),
	}
}
