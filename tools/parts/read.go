package parts

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ReadFile(pfad string) bool {
	fileType := strings.Split(pfad, ".")[len(strings.Split(pfad, "."))-1]
	switch fileType {
	case "xlsx":
		readExcel(pfad)
		return false
	case "json":
		return false
	default:
		return true
	}
}

func readExcel(pfad string) {
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
	//for a, b := range headers {fmt.Printf("Header: %-25s Column: %-25d \n", a, b)}
	headersClean, err := headersClean(headers)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Headers Clean")
	//for a, b := range headersClean {fmt.Printf("Header: %-25s Column: %-25d \n", a, b)}
	betriebsmittelListe := map[string]*BETRIEBSMITELL{}
	for i := headerRow + 1; i < len(rows)-1; i++ {
		key := ("==" + rows[i][headersClean["FunktionaleZuordnung"]] +
			"=" + rows[i][headersClean["Funktionskennzeichen"]] +
			"++" + rows[i][headersClean["Aufstellungsort"]] +
			"+" + rows[i][headersClean["Ortskennzeichen"]] +
			"-" + rows[i][headersClean["BMK"]])

		_, ok := betriebsmittelListe[key]
		if ok {
			btm := betriebsmittelListe[key]
			btm.Artikel = append(btm.Artikel, NewArtikel(headersClean, rows[i]))
		} else {
			btm := NewBetriebsmittel(headersClean, rows[i])
			btm.Artikel = append(btm.Artikel, NewArtikel(headersClean, rows[i]))
			betriebsmittelListe[key] = btm
		}
	}
	fmt.Println(betriebsmittelListe)
}

func headersClean(headers map[string]uint64) (map[string]uint64, error) {
	headersClean := map[string]uint64{
		"FunktionaleZuordnung": math.MaxUint64,
		"Funktionskennzeichen": math.MaxUint64,
		"Aufstellungsort":      math.MaxUint64,
		"Ortskennzeichen":      math.MaxUint64,
		"BMK":                  math.MaxUint64,

		"Bestellnummer": math.MaxUint64,
		"ERP_KNT":       math.MaxUint64,
		"Hersteller":    math.MaxUint64,
		"Stueckzahl":    math.MaxUint64,
		"Beschreibung":  math.MaxUint64,
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
