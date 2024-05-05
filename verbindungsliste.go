package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (a *App) VerbindungRead() {

	// Open our xmlFile
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Verbindungsliste.xml")
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
	teilArray := make(map[string]*Verbindung)
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
			
			fmt.Printf("Id: %-20s", line.Labels[b].XMLName.Space)
			fmt.Printf("Id: %-20s", line.Labels[b].Id)
			fmt.Printf("\n")
			teilArray[label.Properties[1].PropertyValue] = &Verbindung{}
			teilArray[label.Properties[1].PropertyValue].SetBetriebsmittel(&line.Labels[b])
		}
	}

	fmt.Println("Verbindungsliste Fertig")
}
func (b *Verbindung) SetBetriebsmittel(L *Label) {
	count := [30]int{}
	for _,prop := range L.Properties{
		switch prop.PropertyName{		
		case "ID":  
			b.ID = prop.PropertyValue
		case "Name des Zielanschlusses (vollständig)":
			if count[0] == 0{
				b.Bauteil[0].BMK.BMKVoll = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.BMKVoll = prop.PropertyValue
			}		
			count[0]++	
		case "BMK (identifizierend)":   
			if count[1] == 0{
				b.Bauteil[0].BMK.BMKID = prop.PropertyValue
			}else if count[1] == 1{
				b.Bauteil[1].BMK.BMKID = prop.PropertyValue
			}else{
				b.Name.BMKID = prop.PropertyValue
			}
			count[1]++	
		case "Funktionale Zuordnung":                             
			if count[2] == 0{
				b.Bauteil[0].BMK.FunktionaleZuordnung = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.FunktionaleZuordnung = prop.PropertyValue
			}
			count[2]++	  
		case "Funktionskennzeichen":                              
			if count[3] == 0{
				b.Bauteil[0].BMK.Funktionskennzeichen = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.Funktionskennzeichen = prop.PropertyValue
			}   
			count[3]++	
		case "Aufstellungsort":                                  
 			if count[4] == 0{
				b.Bauteil[0].BMK.Aufstellungsort = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.Aufstellungsort = prop.PropertyValue
			}
			count[4]++	   
		case "Ortskennzeichen":                                   
 			if count[5] == 0{
				b.Bauteil[0].BMK.Ortskennzeichen = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.Ortskennzeichen = prop.PropertyValue
			}   
			count[5]++	
		case "BMK (identifizierend, ohne Projektstrukturen)":
 			if count[6] == 0{
				b.Bauteil[0].BMK.BMK = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.BMK = prop.PropertyValue
			} 
			count[6]++	  
		case "BMK: Kennbuchstabe":                                
 			if count[7] == 0{
				b.Bauteil[0].BMK.BMKKennbuchstabe = prop.PropertyValue
			}else{
				b.Bauteil[1].BMK.BMKKennbuchstabe = prop.PropertyValue
			} 
			count[7]++	  
		case "Funktionstext":                                    
				b.Bauteil[1].Funktionstext = prop.PropertyValue  
		case "Technische Kenngrößen":                            
 			if count[8] == 0{
				b.Bauteil[0].TechnischeKenngrößen = prop.PropertyValue
			}else{
				
				b.Bauteil[1].TechnischeKenngrößen = prop.PropertyValue
			}  
			count[8]++	 
		case "Funktionsdefinition: Kategorie":                    
 			if count[9] == 0{
				b.Bauteil[0].FunktionsdefinitionKategorie  = prop.PropertyValue
			}else{
				b.Bauteil[1].FunktionsdefinitionKategorie = prop.PropertyValue
			} 
			count[9]++	  
		case "Funktionsdefinition: Gruppe":                       
 			if count[10] == 0{
				b.Bauteil[0].FunktionsdefinitionGruppe = prop.PropertyValue
			}else{
				b.Bauteil[1].FunktionsdefinitionGruppe = prop.PropertyValue
			} 
			count[10]++	  
		case "Funktionsdefinition: Beschreibung":                 
 			if count[11] == 0{
				b.Bauteil[0].FunktionsdefinitionBeschreibung = prop.PropertyValue
			}else{
				b.Bauteil[1].FunktionsdefinitionBeschreibung = prop.PropertyValue
			} 
			count[11]++	  
		case "Anschlussbezeichnung der Funktion":                 
 			if count[12] == 0{
				b.Bauteil[0].AnschlussbezeichnungderFunktion = prop.PropertyValue
			}else{
				b.Bauteil[1].AnschlussbezeichnungderFunktion = prop.PropertyValue
			} 
			count[12]++	  
		case "Funktionsdefinition":                               
 			if count[13] == 0{
				b.Bauteil[0].Funktionsdefinition = prop.PropertyValue
			}else{
				b.Bauteil[1].Funktionsdefinition = prop.PropertyValue
			}   
			count[13]++	
		case "Symbolname":                                        
 			if count[14] == 0{
				b.Bauteil[0].Symbolname = prop.PropertyValue
			}else{
				b.Bauteil[1].Symbolname = prop.PropertyValue
			}   
			count[14]++	
		case "Symbolvariante":                                    			                          
 			if count[15] == 0{
				b.Bauteil[0].Symbolvariante  = prop.PropertyValue
			}else{
				b.Bauteil[1].Symbolvariante = prop.PropertyValue
			}   
			count[15]++	
		case "Verbindung: Zugehörigkeit":                         			                            
				b.VerbindungZugehörigkeit = prop.PropertyValue  
		case "Verbindungsquerschnitt / -durchmesser":             
				b.Verbindungsquerschnitt = prop.PropertyValue 
		case "Verbindungsfarbe / -nummer":                        
				b.Verbindungsfarbeundnummer = prop.PropertyValue
		case "Verbindung: Länge (vollständig)": 
				i, err := strconv.Atoi(prop.PropertyValue)
				if err != nil {
					fmt.Println(err)
				}                 
				b.VerbindungLänge = i
		case "Netzname":                                         
				b.Netzname = prop.PropertyValue 
		case "Signalname":                                       
				b.Signalname = prop.PropertyValue
		case "Potenzialname":                                     
				b.Potenzialname = prop.PropertyValue
		case "Potenzialtyp":                                      
				b.Potenzialtyp = prop.PropertyValue
		case "Potenzialwert":                                    
				b.Potenzialwert = prop.PropertyValue
		case "Netzindex":
				b.Netzindex = prop.PropertyValue
		default:
			fmt.Printf("Property Name: %-50s", prop.PropertyName)
			fmt.Printf("Property Value: %-50s", prop.PropertyValue)
			fmt.Printf("\n")
		}
	
	}
}