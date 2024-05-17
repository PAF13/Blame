package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func writeCSVFile(fileName string, rows *[][]string, ziel string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	csvFile, err := os.Create(rootPfadOutput + "Blame_" + fileName + ".csv")
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
func (b *VERBINDUNG) writeCSVFileBuffer(rows *[][]string) {
	array := [21]string{}
	array[0] = b.Quelle.BMKVollständig
	array[1] = ""
	array[2] = ""
	array[3] = ""
	array[4] = b.Ziel.BMKVollständig
	array[5] = ""
	array[6] = ""
	array[7] = ""
	array[8] = strings.Replace(fmt.Sprintf("%.1f", b.Verbindungsquerschnitt), ".", ",", 1)
	array[9] = b.Verbindungsfarbeundnummer
	array[10] = fmt.Sprintf("%d", b.VerbindungLänge)
	array[11] = ""
	array[12] = ""
	array[13] = ""
	array[14] = ""
	array[15] = ""
	array[16] = ""
	array[17] = ""
	array[18] = ""
	array[19] = ""
	array[20] = ""
	*rows = append(*rows, array[:])
}
func (b *BETRIEBSMITELL) writeCSVFileBuffer(rows *[][]string) {
	array := [110]string{}
	array[0] = b.BMK.BMKVollständig
	array[1] = b.BMK.BMKVollständig
	array[2] = ""
	array[3] = ""
	array[4] = ""
	array[5] = ""
	array[6] = ""
	array[7] = ""
	array[8] = ""
	array[9] = ""
	index := 10
	for i := 0; i < 50; i++ {
		array[index] = b.Artikel[i].Bestellnummer
		index++
		array[index] = b.Artikel[i].ArtikelnummerEplan
		index++
	}

	*rows = append(*rows, array[:])
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
