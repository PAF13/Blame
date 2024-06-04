package parts

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeCSVFile(fileName string, rows *[][]string, ziel string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	csvFile, err := os.Create("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\" + fileName + ".csv")
	if err != nil {
		log.Printf("failed creating file: %s\n", err)
	}
	defer csvFile.Close()
	w := csv.NewWriter_REFAC(csvFile)
	defer w.Flush()
	writeCSVFileHeader(w, ziel)
	writeCSVFileBody(w, rows)

}
func writeCSVFileBody(w *csv.Writer, rows *[][]string) {
	for _, row := range *rows {
		if err := w.Write(row); err != nil {
			log.Println("error writing record to file", err)
		}
	}

}

func writeCSVFileHeader(w *csv.Writer, ziel string) {
	var header *[]string

	switch {
	case ziel == "PWA6000":
		header = &[]string{
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
			"Quelle Verlegerichtung",
			"Ziel Verlegerichtung",
			"Source Parent Item",
			"Destination Parent Item",
			"Source Z Pos",
			"Destination Z Pos",
		}
	case ziel == "stuecklisteSum":
		header = &[]string{
			"",
			"",
			"",
			"Stuecklistenart",
			"",
			"Pos.",
			"ERP",
			"Menge",
			"Herstellernummer",
			"Hersteller",
			"Quelle",
		}
	case ziel == "EPlanBetriebsmittel":
		header = &[]string{
			"BMK Vollständig",
			"BMK Id.",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
		}
		for i := 1; i < 51; i++ {
			*header = append(*header,
				"Bestellnummer["+fmt.Sprintf("%d", i)+"]",
				"Artikelnummer["+fmt.Sprintf("%d", i)+"]",
			)
		}
	case ziel == "noHeader":
		header = &[]string{}
	}

	if err := w.Write(*header); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}
