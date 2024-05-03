package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"runtime"
)

/*
Name des Zielanschlusses (vollständig)
BMK (identifizierend)
Funktionale Zuordnung
Funktionskennzeichen
Aufstellungsort
Ortskennzeichen
BMK (identifizierend, ohne Projektstrukturen)
BMK: Kennbuchstabe
Funktionstext
Technische Kenngrößen
Funktionsdefinition: Kategorie
Funktionsdefinition: Gruppe
Funktionsdefinition: Beschreibung
Anschlussbezeichnung der Funktion
Funktionsdefinition
Symbolname
Symbolvariante
*/

func (a *App) VerbindungRead(pfad string) {

	// Open our xmlFile
	xmlFile, err := os.Open(pfad)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened EPlan_Klemmen.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)

	// we initialize our Users array
	var eplanLabelling EplanLabelling
	teilArray := make(map[string]*Betriebsmittel)
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &eplanLabelling)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for a := 0; a < len(eplanLabelling.Document.Page.Lines); a++ {
		line := eplanLabelling.Document.Page.Lines[a]
		for b := 0; b < len(line.Labels); b++ {
			label := eplanLabelling.Document.Page.Lines[a].Labels[b]
			fmt.Printf("Id: %-20s", line.Labels[b].Id)
			fmt.Printf("BMK: %-50s", line.Labels[b].Properties[1].PropertyValue)
			fmt.Printf("\n")
			teilArray[label.Properties[1].PropertyValue] = &Betriebsmittel{}
			teilArray[label.Properties[1].PropertyValue].SetBetriebsmittel(&line.Labels[b])
		}
	}

	fmt.Println("Verbindungsliste Fertig")
}

func (b *Betriebsmittel) SetBetriebsmittel(L *Label) {
	for c := 0; c < len(L.Properties); c++ {
		switch os := runtime.GOOS; os {
		case "BMK (identifizierend)":
			b.BMKVollständig = L.Properties[c].PropertyValue
		case "Funktionale Zuordnung":
			b.FunktionaleZuordnung = L.Properties[c].PropertyValue
		case "Funktionskennzeichen":
			b.Funktionskennzeichen = L.Properties[c].PropertyValue
		case "Aufstellungsort":
			b.Aufstellungsort = L.Properties[c].PropertyValue
		case "Ortskennzeichen":
			b.Ortskennzeichen = L.Properties[c].PropertyValue
		//case "linux":
		//Dokumentenart = L.Properties[C].PropertyValue
		//case "linux":
		//BenutzerdefinierteStruktur = L.Properties[C].PropertyValue
		//case "linux":
		//Anlagennummer = L.Properties[C].PropertyValue
		case "linux":
			b.BMK = L.Properties[c].PropertyValue
		//case "Artikelnummer":
		//b.Artikel = append(b.Artikel, L.Properties[c].PropertyValue)
		default:
			fmt.Printf("Property Name: %-50s", L.Properties[c].PropertyName)
			fmt.Printf("Property Value: %-50s", L.Properties[c].PropertyValue)
			fmt.Printf("\n")
		}
	}
}
