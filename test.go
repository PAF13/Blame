package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
	"golang.org/x/text/unicode/norm"
)

/*
structure is as fallows:
- Parts will be under a interface to take care of
  - internal connection detection of anschlusse ie.  L1 = U = T1 etc
*/
type Verbindung struct {
	ID              int
	verbindungArt   string
	zugehörigkeit   string
	color           string
	length          string
	querschnitt     string
	ort             [2]string
	anschluss       [2]string
	anschlussSymbol [2]string
}

func (a *App) VerbindungCorrection() {
	start := time.Now()

	verbindungen := []Verbindung{}
	bauteile := make(map[string]string)
	pfad := a.Dialog()
	loadVerbindungsliste(bauteile, &verbindungen, pfad)

	writeVerbindungsliste(&verbindungen)

	//load data into structures
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

func loadVerbindungsliste(bauteile map[string]string, verbindungen *[]Verbindung, pfad string) {

	spreadsheet, err := excelize.OpenFile(pfad)
	STDerrhandler(err)
	defer func() {
		// Close the spreadsheet.
		if err := spreadsheet.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := spreadsheet.Rows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	rows2, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	STDerrhandler(err)
	rowlen := len(rows2)
	i := 0
	for rows.Next() {
		row, err := rows.Columns()
		STDerrhandler(err)

		verbindungCLean := Verbindung{
			ID:              i,
			verbindungArt:   safeStringArrayPull(row, 36),
			zugehörigkeit:   safeStringArrayPull(row, 37),
			color:           safeStringArrayPull(row, 38),
			length:          safeStringArrayPull(row, 39),
			querschnitt:     safeStringArrayPull(row, 37),
			anschluss:       [2]string{safeStringArrayPull(row, 0), safeStringArrayPull(row, 17)},
			anschlussSymbol: [2]string{safeStringArrayPull(row, 15), safeStringArrayPull(row, 32)},
			ort:             [2]string{safeStringArrayPull(row, 5), safeStringArrayPull(row, 22)}}

		*verbindungen = append(*verbindungen, verbindungCLean)

		_, ok1 := bauteile[norm.NFD.String(safeStringArrayPull(row, 1))]
		if !ok1 {
			bauteile[norm.NFD.String(safeStringArrayPull(row, 1))] = ""
		}

		_, ok2 := bauteile[norm.NFD.String(safeStringArrayPull(row, 18))]
		if !ok2 {
			bauteile[norm.NFD.String(safeStringArrayPull(row, 18))] = ""
		}
		//		if !ok1 && !strings.Contains(norm.NFD.String(safeStringArrayPull(row, 1)), norm.NFD.String("_")) && !strings.Contains(norm.NFD.String(safeStringArrayPull(row, 1)), norm.NFD.String("-")) {
		//		if !ok2 && !strings.Contains(norm.NFD.String(safeStringArrayPull(row, 18)), norm.NFD.String("_")) && !strings.Contains(norm.NFD.String(safeStringArrayPull(row, 18)), norm.NFD.String("-")) {

		i++
	}

	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}

	for a, b := range bauteile {
		fmt.Println(a, b)
	}
	//Load lists: Verbindungsliste und Betriebsmittelliste
	fmt.Println("Row #: ", rowlen)
}

func writeVerbindungsliste(verbindungen *[]Verbindung) {
	file, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Verbindungsliste_WIR.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.Comma = ';'
	defer w.Flush()

	file2, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Verbindungsliste_Fehlerhaft_WIR.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file2.Close()
	w2 := csv.NewWriter(file2)
	w2.Comma = ';'
	defer w2.Flush()

	// Using Write
	verbindungslisteHeader := []string{
		"DEST1",
		"TA1",
		"STOP1",
		"DEVICE1",
		"DEST2",
		"TA2",
		"STOP2",
		"DEVICE2",
		"CROSSSECTION",
		"COLOUR",
		"LENGTH",
		"TYPE",
		"WIREID",
		"CAEID",
		"DIR1",
		"DIR2",
		"PARENT1",
		"PARENT2",
		"Z1",
		"Z2"}
	if err := w.Write(verbindungslisteHeader); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	if err := w2.Write(verbindungslisteHeader); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	for _, record := range *verbindungen {
		if record.ort != [2]string{"", ""} || record.color != "" || record.querschnitt != "" {
			row := []string{
				record.anschluss[0],
				"", //TA1
				"3",
				"", //DEVICE1
				record.anschluss[1],
				"", //TA2
				"3",
				"", //DEVICE2
				record.querschnitt,
				record.color,
				record.length,
				"", //TYPE
				"", //WIREID
				"", //CAEID
				"", //DIR1
				"", //DIR2
				"", //PARENT1
				"", //PARENT2
				"", //Z1
				"" /*Z2*/}
			if record.length == "" {
				if err := w2.Write(row); err != nil {
					log.Fatalln("error writing record to file", err)
				}
			} else {
				if err := w.Write(row); err != nil {
					log.Fatalln("error writing record to file", err)
				}
			}
		}
	}

}
