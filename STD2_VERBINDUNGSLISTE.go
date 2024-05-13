package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verbindung_Clean []VERBINDUNG

func INIT_VERBINDUNGSLITE() {
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Verbindungsliste.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	verbindung_Clean = []VERBINDUNG{}

	b2, err := json.MarshalIndent(verbindung_Clean, "", "    ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\blame_verbindungsliste.json", b2, 0644)
	if err != nil {
		log.Println(err)
	}

	STD_Write_Verbindungsliste(verbindung_Clean)
	fmt.Println("Verbindungsliste Fertig")
}

/*
	func (structType *EplanAuswertungXML) convertFile(byteValue []byte) {
		verbindungsliste := map[string]VERBINDUNG{}
		xml.Unmarshal(byteValue, &structType)
		line := structType.Document.Page.Line
		for a, aa := range line {
			fmt.Printf("Source ID: %-50s", aa.Label.SourceID)
			fmt.Printf("Anzahl: %-50d\n", a+1)
			verbindung := VERBINDUNG{}
			P_verbindung := &verbindung
			for _, bb := range aa.Label.Property {

				switch bb.PropertyName {
				case "Name des Zielanschlusses (vollständig)":
					if P_verbindung.Quelle.BMKVollständig != "" {
						P_verbindung.Ziel.BMKVollständig = bb.PropertyValue
					} else {
						P_verbindung.Quelle.BMKVollständig = bb.PropertyValue
					}
				case "BMK (identifizierend)":
					if P_verbindung.Quelle.BMKidentifizierung != "" {
						P_verbindung.Ziel.BMKidentifizierung = bb.PropertyValue
					} else {
						P_verbindung.Quelle.BMKidentifizierung = bb.PropertyValue
					}
				case "Funktionale Zuordnung":
					if P_verbindung.Quelle.FunktionaleZuordnung != "" {
						P_verbindung.Ziel.FunktionaleZuordnung = bb.PropertyValue
					} else {
						P_verbindung.Quelle.FunktionaleZuordnung = bb.PropertyValue
					}
				case "Funktionskennzeichen":
					if P_verbindung.Quelle.Funktionskennzeichen != "" {
						P_verbindung.Ziel.Funktionskennzeichen = bb.PropertyValue
					} else {
						P_verbindung.Quelle.Funktionskennzeichen = bb.PropertyValue
					}
				case "Aufstellungsort":
					if P_verbindung.Quelle.Aufstellungsort != "" {
						P_verbindung.Ziel.Aufstellungsort = bb.PropertyValue
					} else {
						P_verbindung.Quelle.Aufstellungsort = bb.PropertyValue
					}
				case "Ortskennzeichen":
					if P_verbindung.Quelle.Ortskennzeichen != "" {
						P_verbindung.Ziel.Ortskennzeichen = bb.PropertyValue
					} else {
						P_verbindung.Quelle.Ortskennzeichen = bb.PropertyValue
					}
				case "BMK (identifizierend, ohne Projektstrukturen)":
					if P_verbindung.Quelle.BMK != "" {
						P_verbindung.Ziel.BMK = bb.PropertyValue
					} else {
						P_verbindung.Quelle.BMK = bb.PropertyValue
					}
				case "BMK: Kennbuchstabe":
					if P_verbindung.Quelle.Kennbuchstabe != "" {
						P_verbindung.Ziel.Kennbuchstabe = bb.PropertyValue
					} else {
						P_verbindung.Quelle.Kennbuchstabe = bb.PropertyValue
					}

				case "Verbindung: Zugehörigkeit":
					P_verbindung.VerbindungZugehörigkeit = bb.PropertyValue
				case "Verbindungsquerschnitt / -durchmesser":
					P_verbindung.Verbindungsquerschnitt, _ = strconv.ParseFloat(bb.PropertyValue, 64)
				case "Verbindungsfarbe / -nummer":
					P_verbindung.Verbindungsfarbeundnummer = bb.PropertyValue
				case "Verbindung: Länge (vollständig)":
					P_verbindung.VerbindungLänge, _ = strconv.Atoi(bb.PropertyValue)
				case "Netzname":
					P_verbindung.Netzname = bb.PropertyValue
				case "Signalname":
					P_verbindung.Signalname = bb.PropertyValue
				case "Potenzialname":
					P_verbindung.Potenzialname = bb.PropertyValue
				case "Potenzialtyp":
					P_verbindung.Potenzialtyp = bb.PropertyValue
				case "Potenzialwert":
					P_verbindung.Potenzialwert = bb.PropertyValue
				case "Netzindex":
					P_verbindung.Netzindex = bb.PropertyValue
				default:
					fmt.Printf("Missing | Name: %-50s", bb.PropertyName)
					fmt.Printf("Value: %-50s", bb.PropertyValue)
					fmt.Printf("\n")
				}

			}

			verbindungsliste[aa.Label.SourceID] = *P_verbindung
		}

		content, err := json.MarshalIndent(verbindungsliste, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\blame_verbindungsliste2.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
*/
func STD_Write_Verbindungsliste(bind []VERBINDUNG) {
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
