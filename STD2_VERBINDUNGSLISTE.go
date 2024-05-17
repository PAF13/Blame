package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func (a *App) ImportVerbindungsliste() {
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Verbindungsliste.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	XML := NewEplanAuswertungXML()

	xml.Unmarshal(byteValue, XML)

	XML.STD_Write_Verbindungsliste(byteValue)
	fmt.Println("Verbindungsliste Fertig")
}
func setVerbindung(quelle *string, ziel *string, pos *bool, value string) {

	if !*pos {
		*quelle = value
		*pos = true
	} else {
		*ziel = value
	}
}
func (structType *EplanAuswertungXML) STD_Write_Verbindungsliste(byteValue []byte) {
	verbindungsliste := NewVerbindungMap()
	array := NewVerbindungArray()
	verbindungAllpolig := NewVerbindungArray()
	verbindung3D := NewVerbindungArray()
	xml.Unmarshal(byteValue, &structType)
	line := structType.Document.Page.Line
	for a, aa := range line {
		fmt.Println("line: " + fmt.Sprintf("%d", a))
		P_verbindung := NewVerbindung()
		pos := make(map[string]*bool)
		for _, bb := range aa.Label.Property {
			switch bb.PropertyName {
			case "Name des Zielanschlusses (vollständig)":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.BMKVollständig), &(P_verbindung.Ziel.BMKVollständig), pos[bb.PropertyName], bb.PropertyValue)
			case "BMK (identifizierend)":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.BMKidentifizierung), &(P_verbindung.Ziel.BMKidentifizierung), pos[bb.PropertyName], bb.PropertyValue)
			case "Funktionale Zuordnung":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.FunktionaleZuordnung), &(P_verbindung.Ziel.FunktionaleZuordnung), pos[bb.PropertyName], bb.PropertyValue)
			case "Funktionskennzeichen":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Funktionskennzeichen), &(P_verbindung.Ziel.Funktionskennzeichen), pos[bb.PropertyName], bb.PropertyValue)
			case "Aufstellungsort":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Aufstellungsort), &(P_verbindung.Ziel.Aufstellungsort), pos[bb.PropertyName], bb.PropertyValue)
			case "Ortskennzeichen":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Ortskennzeichen), &(P_verbindung.Ziel.Ortskennzeichen), pos[bb.PropertyName], bb.PropertyValue)
			case "BMK (identifizierend, ohne Projektstrukturen)":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.BMK), &(P_verbindung.Ziel.BMK), pos[bb.PropertyName], bb.PropertyValue)
			case "BMK: Kennbuchstabe":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Kennbuchstabe), &(P_verbindung.Ziel.Kennbuchstabe), pos[bb.PropertyName], bb.PropertyValue)
			case "Symbolname":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Symbolname), &(P_verbindung.Ziel.Symbolname), pos[bb.PropertyName], bb.PropertyValue)
			case "Platzierung":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.Quelle.Platzierung), &(P_verbindung.Ziel.Platzierung), pos[bb.PropertyName], bb.PropertyValue)
			case "Anschlussbezeichnung der Funktion":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.QuelleAnschlussbezeichnungderFunktion), &(P_verbindung.ZielAnschlussbezeichnungderFunktion), pos[bb.PropertyName], bb.PropertyValue)
			case "Klemmen- / Steckerkontaktbezeichnung":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.QuelleKlemmenbezeichnung), &(P_verbindung.ZielKlemmenbezeichnung), pos[bb.PropertyName], bb.PropertyValue)
			case "Funktionsdefinition: Kategorie":
				_, ok := pos[bb.PropertyName]
				if !ok {
					pos[bb.PropertyName] = NewBool()
				}
				setVerbindung(&(P_verbindung.QuelleFunktionKatagorie), &(P_verbindung.ZielFunktionKatagorie), pos[bb.PropertyName], bb.PropertyValue)
			case "Verbindung: Zugehörigkeit":
				P_verbindung.VerbindungZugehörigkeit = bb.PropertyValue
			case "Verbindungsquerschnitt / -durchmesser":
				P_verbindung.Verbindungsquerschnitt, _ = strconv.ParseFloat(strings.Replace(bb.PropertyValue, ",", ".", 1), 64)
			case "Verbindungsfarbe / -nummer":
				P_verbindung.Verbindungsfarbeundnummer = bb.PropertyValue
			case "Verbindung: Länge (vollständig)":
				laenge, _ := strconv.Atoi(bb.PropertyValue)

				P_verbindung.VerbindungLänge = laenge

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
			case "Funktionsdefinition":
				P_verbindung.Funktionsdefinition = bb.PropertyValue
			case "Darstellungsart":
				P_verbindung.Darstellungsart = bb.PropertyValue

			default:
				//fmt.Printf("Missing | Name: %-50s", bb.PropertyName)
				//fmt.Printf("Value: %-50s", bb.PropertyValue)
				//fmt.Printf("\n")
			}

		}
		/*ort := "P2_CS1"
		if P_verbindung.Ziel.Ortskennzeichen == ort || P_verbindung.Quelle.Ortskennzeichen == ort {
			_, ok := verbindungsliste[SortString(P_verbindung.Ziel.BMKVollständig+P_verbindung.Quelle.BMKVollständig)]
			if ok {
				fmt.Println(P_verbindung.Ziel.BMKVollständig + " to " + P_verbindung.Quelle.BMKVollständig + " was not added")
			}


		}*/
		if P_verbindung.QuelleKlemmenbezeichnung == "" {
			P_verbindung.QuelleAnschluss = P_verbindung.QuelleAnschlussbezeichnungderFunktion
		} else {
			P_verbindung.QuelleAnschluss = P_verbindung.QuelleKlemmenbezeichnung
		}

		if P_verbindung.ZielKlemmenbezeichnung == "" {
			P_verbindung.ZielAnschluss = P_verbindung.ZielAnschlussbezeichnungderFunktion
		} else {
			P_verbindung.ZielAnschluss = P_verbindung.ZielKlemmenbezeichnung
		}

		*array = append(*array, *P_verbindung)
		if P_verbindung.Verbindungsquerschnitt <= 6 && P_verbindung.VerbindungZugehörigkeit == "Einzelverbindung" && P_verbindung.Funktionsdefinition == "Ader / Draht" {
			switch {
			case P_verbindung.Darstellungsart == "Allpolig":
				*verbindungAllpolig = append(*verbindungAllpolig, *P_verbindung)
			case P_verbindung.Darstellungsart == "3D-Montageaufbau":

				*verbindung3D = append(*verbindung3D, *P_verbindung)
			}
			// P_verbindung.Verbindungsquerschnitt <= 6 && P_verbindung.VerbindungZugehörigkeit == "Einzelverbindung" && P_verbindung.Funktionsdefinition != "Stegbrücke"

		}

	}
	//fmt.Println(verbindung3D)
	for _, b3 := range *verbindungAllpolig {
		_, ok := verbindungsliste[b3.Quelle.BMKVollständig+b3.Ziel.BMKVollständig]
		if !ok {
			verbindungsliste[b3.Quelle.BMKVollständig+b3.Ziel.BMKVollständig] = writeVerbindung(b3)
		}
	}
	for _, b2 := range *verbindung3D {
		_, ok := verbindungsliste[b2.Quelle.BMKVollständig+b2.Ziel.BMKVollständig]
		if !ok {
			verbindungsliste[b2.Quelle.BMKVollständig+b2.Ziel.BMKVollständig] = writeVerbindung(b2)
		} else {
			if b2.VerbindungLänge != 0 {
				verbindungsliste[b2.Quelle.BMKVollständig+b2.Ziel.BMKVollständig] = updateVerbindung(verbindungsliste[b2.Quelle.BMKVollständig+b2.Ziel.BMKVollständig], b2.VerbindungLänge)
			}

		}
	}

	content, err := json.MarshalIndent(verbindungsliste, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\blame_verbindungsliste2.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
	STD_Write_Verbindungsliste(verbindungsliste, "map")
	STD_Write_Verbindungsliste2(*array, "array")
}
func updateVerbindung(v VERBINDUNG, l int) VERBINDUNG {
	v.VerbindungLänge = l
	return v
}
func writeVerbindung(v VERBINDUNG) VERBINDUNG {
	if v.VerbindungLänge < 300 && v.VerbindungLänge != 0 {
		v.VerbindungLänge = 300
	}
	return v
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func STD_Write_Verbindungsliste(bind map[string]VERBINDUNG, typeName string) {
	csvFile, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Verbindungsliste_" + typeName + ".csv")
	if err != nil {
		log.Printf("failed creating file: %s\n", err)
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

		verbindung[0] = b.Quelle.BMKVollständig
		verbindung[1] = ""
		verbindung[2] = "3"
		verbindung[3] = ""
		verbindung[4] = b.Ziel.BMKVollständig
		verbindung[5] = ""
		verbindung[6] = "3"
		verbindung[7] = ""
		//verbindung[8] = fmt.Sprintf("%.1f", b.Verbindungsquerschnitt)
		verbindung[8] = strings.Replace(fmt.Sprintf("%.1f", b.Verbindungsquerschnitt), ".", ",", 1)
		verbindung[9] = b.Verbindungsfarbeundnummer
		verbindung[10] = fmt.Sprintf("%d", b.VerbindungLänge)
		verbindung[11] = ""
		verbindung[12] = b.ZielFunktionKatagorie
		verbindung[13] = b.VerbindungZugehörigkeit
		verbindung[14] = b.Funktionsdefinition
		verbindung[15] = b.Darstellungsart
		if err := w.Write(verbindung); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}

func STD_Write_Verbindungsliste2(bind []VERBINDUNG, typeName string) {
	csvFile, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Verbindungsliste_" + typeName + ".csv")
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

		verbindung[0] = b.Quelle.BMKVollständig
		verbindung[1] = b.QuelleAnschluss
		verbindung[2] = "3"
		verbindung[3] = ""
		verbindung[4] = b.Ziel.BMKVollständig
		verbindung[5] = b.ZielAnschluss
		verbindung[6] = "3"
		verbindung[7] = ""
		verbindung[8] = strings.Replace(fmt.Sprintf("%.1f", b.Verbindungsquerschnitt), ".", ",", 1)
		verbindung[9] = b.Verbindungsfarbeundnummer
		verbindung[10] = fmt.Sprintf("%d", b.VerbindungLänge)
		verbindung[11] = ""
		verbindung[12] = b.ZielFunktionKatagorie
		verbindung[13] = b.VerbindungZugehörigkeit
		verbindung[14] = b.Funktionsdefinition
		verbindung[15] = b.Darstellungsart
		if err := w.Write(verbindung); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
