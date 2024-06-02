package parts

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/xuri/excelize/v2"
)

func LoadStueckliste(pfad []string) {
	betriebsmittelListe := NewBetriebsmillliste()
	betriebsmittelListe.importListe()

}

func LoadLager(pfaden []string) {
	lagerliste := NewLagerliste()
	for _, b := range pfaden {
		lagerliste.importListe(b)
	}

}

func (Liste *LAGERLISTE) importListe(pfad string) {
	fileName := strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0]
	fmt.Println(fileName)
	rows := readExcel(pfad)
	setHeader(rows)
}

func (Liste *BETRIEBSMITELLLISTE) importListe() {

}
func readExcel(pfad string) [][]string {
	file, err := excelize.OpenFile(pfad)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func setHeader(rows [][]string) {
	headers := make(map[string]uint64)

	for i := 0; i < 20; i++ {
		for a, b := range rows[i] {
			headers[stringCleaner(b)] = uint64(a)
		}
		headersClean, err := headersClean(headers)
		if err != nil {
			fmt.Println(err)
		}

		if (headersClean["Bestellnummer"] != math.MaxUint64 &&
			headersClean["Hersteller"] != math.MaxUint64 &&
			headersClean["Stueckzahl"] != math.MaxUint64) &&
			(headersClean["ERP"] != math.MaxUint64 ||
				headersClean["ERP_KNT"] != math.MaxUint64) {

		}
	}
}
func headersClean(headers map[string]uint64) (map[string]uint64, error) {
	headersClean := map[string]uint64{
		"FunktionaleZuordnung": math.MaxUint64,
		"Funktionskennzeichen": math.MaxUint64,
		"Aufstellungsort":      math.MaxUint64,
		"Ortskennzeichen":      math.MaxUint64,
		"BMK":                  math.MaxUint64,

		"Bestellnummer": math.MaxUint64,
		"ERP":           math.MaxUint64,
		"ERP_KNT":       math.MaxUint64,
		"Hersteller":    math.MaxUint64,
		"Stueckzahl":    math.MaxUint64,
		"Beschreibung":  math.MaxUint64,

		"Beistellung": math.MaxUint64,
	}
	translate := map[string]string{
		"MODUL==":       "FunktionaleZuordnung",
		"FUNKTION=":     "Funktionskennzeichen",
		"AUFSTELLORT++": "Aufstellungsort",
		"EINBAUORT+":    "Ortskennzeichen",
		"BMK-":          "BMK",

		"BESTELLNUMMER": "Bestellnummer",
		"ERP-NUMMER":    "ERP_KNT",
		"HERSTELLER":    "Hersteller",
		"STUECKZAHL":    "Stueckzahl",
		"BEZEICHNUNG":   "Beschreibung",

		"TEILEART": "Beistellung",
	}

	for a, b := range headers {
		_, ok := translate[a]
		if ok {
			headersClean[translate[a]] = b
		} else {
			fmt.Printf("Header %-25s not found\n", a)
		}
	}
	var err error
	for a, b := range headersClean {
		if b == math.MaxUint64 {
			err = New("missing: " + a)
		}
	}
	return headersClean, err
}
func readExcel1(pfad string) {
	file, err := excelize.OpenFile(pfad)
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}
	headerRow := 6
	headers := make(map[string]uint64)
	fileLength := len(rows) - headerRow - 1

	fmt.Printf("Header Row: %-25d File Length: %-25d \n", headerRow, fileLength)
	for a, b := range rows[headerRow] {
		headers[stringCleaner(b)] = uint64(a)
	}
	for a, b := range headers {
		fmt.Printf("Header: %-25s Column: %-25d \n", a, b)
	}
	headersClean, err := headersClean(headers)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Headers Clean")
	for a, b := range headersClean {
		fmt.Printf("Header: %-25s Column: %-25d \n", a, b)
	}

	betriebsmittelListe := BETRIEBSMITELLLISTE{
		Betriebsmittel: make(map[string]*BETRIEBSMITELL),
	}
	ortListe := FILTER{
		Filter: make(map[string]bool),
	}
	for i := headerRow + 1; i < len(rows)-1; i++ {
		keyBetriebsmittel := ("==" + rows[i][headersClean["FunktionaleZuordnung"]] +
			"=" + rows[i][headersClean["Funktionskennzeichen"]] +
			"++" + rows[i][headersClean["Aufstellungsort"]] +
			"+" + rows[i][headersClean["Ortskennzeichen"]] +
			"-" + rows[i][headersClean["BMK"]])

		_, okBetriebsmittel := betriebsmittelListe.Betriebsmittel[keyBetriebsmittel]
		if okBetriebsmittel {
			btm := betriebsmittelListe.Betriebsmittel[keyBetriebsmittel]
			btm.Artikel = append(btm.Artikel, NewArtikelTemp(headersClean, rows[i]))
		} else {
			btm := NewBetriebsmittelTemp(headersClean, rows[i])
			btm.Artikel = append(btm.Artikel, NewArtikelTemp(headersClean, rows[i]))
			betriebsmittelListe.Betriebsmittel[keyBetriebsmittel] = btm
		}
		keyOrt := rows[i][headersClean["Ortskennzeichen"]]
		bestellung := strings.ToUpper(rows[i][headersClean["Beistellung"]])
		_, okOrt := ortListe.Filter[keyOrt]

		if okOrt && bestellung == "SITECA" {
			ortListe.Filter[keyOrt] = true
		} else if !okOrt && bestellung == "SITECA" {
			ortListe.Filter[keyOrt] = true
		} else {
			ortListe.Filter[keyOrt] = false
		}
	}

	for a, b := range ortListe.Filter {
		fmt.Printf("Ort: %-25s Bestellung Setica: %-25t \n", a, b)
	}
	writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste_Clean", betriebsmittelListe.Betriebsmittel)
	writeJsonFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste_Clean", betriebsmittelListe)
	writeJsonFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\filter", ortListe)
}

func stringCleaner(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}

func bestellnummerCleaner(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, "-", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "_", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}
