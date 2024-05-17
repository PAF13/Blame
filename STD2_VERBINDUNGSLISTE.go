package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	verbindungtest := &[][]string{}
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
		P_verbindung.writeCSVFileBuffer(verbindungtest)

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

	writeJsonFile("verbindungsliste2", verbindungsliste)
	verbindungtest2 := &[][]string{}
	for _, b := range verbindungsliste {
		b.writeCSVFileBuffer(verbindungtest2)
	}
	writeJsonFile("Verbindungsliste_Raw", verbindungtest)
	writeCSVFile("Verbindungsliste_Raw", verbindungtest, "PWA6000")
	writeJsonFile("Verbindungsliste_Clean", verbindungtest2)
	writeCSVFile("Verbindungsliste_Clean", verbindungtest2, "PWA6000")
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
