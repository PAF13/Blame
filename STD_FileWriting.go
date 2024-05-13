package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func STD_Write_Verbindungsliste(bind []*Verbindung) {
	csvFile, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Verbindungsliste.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	w := csv.NewWriter_REFAC(csvFile)
	defer w.Flush()
	headers := []string{
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

	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	verbindung := []string{}
	for i := 1; i < 20; i++ {
		verbindung = append(verbindung, "")
	}
	for _, b := range bind {

		verbindung[0] = ""
		verbindung[1] = ""
		verbindung[2] = "3"
		verbindung[3] = ""
		verbindung[4] = ""
		verbindung[5] = ""
		verbindung[6] = "3"
		verbindung[7] = ""
		verbindung[8] = strings.Replace(fmt.Sprintf("%.1f", b.Verbindungsquerschnitt), ".", ",", 1)
		verbindung[9] = b.Verbindungsfarbeundnummer
		verbindung[10] = strconv.Itoa(b.VerbindungLänge)
		if err := w.Write(verbindung); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
