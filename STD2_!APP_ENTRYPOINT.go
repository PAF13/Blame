package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/*
	func STD_Write_Stueckliste(Lagerbestand *[]ARTIKEL) {
		file := excelize.NewFile()
		headers := []string{
			"KNT Lager",
		}
		headers2 := []string{
			"ERP",
			"Menge",
			"Hersteller",
			"Bestellnummer",
			"Mehrfach ERP",
			"Beschreibung",
			"Warengruppe",
			"Quelle",
			"Stand",
			"Bereitsteller",
			"Ort",
		}
		for i, header := range headers {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
		}
		for i, header := range headers2 {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 2), header)
		}

		rowNum := 3
		fmt.Println("starting excel")
		for _, value := range *Lagerbestand {
			fmt.Printf("Bestellnummer: %-50s", value.Bestellnummer)
			fmt.Printf("ERP: %-50s", value.ERP)
			fmt.Printf("länge: %-20d", len(value.Fehler))
			fmt.Printf("\n")
			colNum := 0
			erp := ""
			P_erp := &erp
			if len(value.Fehler) > 0 {
				for _, aa := range value.Fehler {
					*P_erp = aa + " | " + *P_erp

				}
			} else {
				*P_erp = value.ERP
			}

			lineWriter(file, "Sheet1", &colNum, &rowNum, erp)
			//lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", value.Stueckzahl))
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Stueckzahl)
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Hersteller)
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnummer)
			lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%d", len(value.Fehler)))
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Beschreibung)
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Quelle)
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
			rowNum++
		}

		if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\blame_KNTLager_Clean.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
*/

func bestellnummerClean(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	return x
}

/*
func bestellnummerClean2(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "-", "")
	x = strings.ReplaceAll(x, "+", "")
	x = strings.ReplaceAll(x, "/", "")
	x = strings.ReplaceAll(x, ",", "")
	return x
}


func (structType *EplanAuswertungXML) convertFile(byteValue []byte) {
	verbindungsliste = map[string]VERBINDUNG{}
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
				P_verbindung.Verbindungsquerschnitt = bb.PropertyValue
			case "Verbindungsfarbe / -nummer":
				P_verbindung.Verbindungsfarbeundnummer = bb.PropertyValue
			case "Verbindung: Länge (vollständig)":
				P_verbindung.VerbindungLänge = bb.PropertyValue
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
	err = ioutil.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\blame_verbindungsliste.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ValueRestrict(s string) error {
	switch s {
	case "Yes", "No", "I don't know":
		fmt.Println("Success!")
		return nil
	default:
		return fmt.Errorf("unsupported value: %q", s)
	}
}
*/
//dialog windows

func (a *App) OpenFileDialog() string {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Projektdaten 2024",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) OpenMultipleFilesDialog() []string {
	result, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Choose File",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Projektdaten 2024",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) OpenDirectoryDialog() string {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) SaveFileDialog() string {
	result, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) MessageDialog() string {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Question",
		Message:       "Do you want to continue?",
		DefaultButton: "No",
	})
	return result
}

// Generate UUID
func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
