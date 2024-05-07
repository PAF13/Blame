package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)



func (a *App) ImportFile(pfad []string) {
	for _,b := range pfad{
		ImportFileIntern(b)
	}
		
}

func ImportFileIntern(pfad string) {
	fileType := ""

		switch fileType{
		case ".xml":
			ImportXML()
		case ".xlsx":
			ImportExcel(pfad)
		case ".json":
			ImportJson(pfad)
		default:

		}
}

func (a *App) ExportFile(pfad []string) {
	for _,b := range pfad{
		ExportFileIntern(b)
	}

}
func ExportFileIntern(pfad string) {
	fileType := ""	
		switch fileType{
		case ".xml":
			ExportXML(pfad)
		case ".xlsx":
			ExportExcel(pfad)
		case ".json":
			ExportJson(pfad)
		default:

		}

}

func  ImportXML() {
// Open our xmlFile
xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Verbindungsliste.xml")
// if we os.Open returns an error then handle it
if err != nil {
	fmt.Println(err)
}

fmt.Println("Successfully Opened users.xml")
// defer the closing of our xmlFile so that we can parse it later on
defer xmlFile.Close()

// read our opened xmlFile as a byte array.
byteValue, _ := io.ReadAll(xmlFile)

// we initialize our Users array
var eplanAuswertungXML EplanAuswertungXML
// we unmarshal our byteArray which contains our
// xmlFiles content into 'users' which we defined above
xml.Unmarshal(byteValue, &eplanAuswertungXML)

// we iterate through every user within our users array and
// print out the user Type, their name, and their facebook url
// as just an example
verbindung := []VERBINDUNG{}
P_verbindung := &verbindung

line := eplanAuswertungXML.Document.Page.Line
for a,aa := range line {
	fmt.Printf("PropertyValue: %-20s", line[a].SourceID)
	*P_verbindung = append(*P_verbindung, VERBINDUNG{
		UUID: line[a].SourceID,
		Quelle: BETRIEBSMITELLKENNZEICHEN{},
		Ziel: BETRIEBSMITELLKENNZEICHEN{},
	})
	fmt.Printf("\n")
		for _,bb := range aa.Label.Property {
			fmt.Printf("PropertyName: %-60s", bb.PropertyName)
			fmt.Printf("PropertyValue: %-20s", bb.PropertyValue)
			fmt.Printf("\n")
		}
	}

	content, err := json.MarshalIndent(verbindung, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Blame_Test.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}


func  ImportExcel(pfad string) {

}

func ImportJson(pfad string) {
// Open our jsonFile
jsonFile, err := os.Open("users.json")
// if we os.Open returns an error then handle it
if err != nil {
	fmt.Println(err)
}

fmt.Println("Successfully Opened users.json")
// defer the closing of our jsonFile so that we can parse it later on
defer jsonFile.Close()

// read our opened xmlFile as a byte array.
byteValue, _ := io.ReadAll(jsonFile)

// we initialize our Users array
var users Settings

// we unmarshal our byteArray which contains our
// jsonFile's content into 'users' which we defined above
json.Unmarshal(byteValue, &users)

// we iterate through every user within our users array and
// print out the user Type, their name, and their facebook url
// as just an example
/*for i := 0; i < len(users.Users); i++ {
	fmt.Println("User Type: " + users.Users[i].Type)
	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	fmt.Println("User Name: " + users.Users[i].Name)
	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
}*/
}

func  ExportXML(pfad string) {

}

func  ExportExcel(pfad string) {

}

func ExportJson(pfad string) {

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

//dialog windows

func (a *App) OpenFileDialog() string {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) OpenMultipleFilesDialog() []string {
	result, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
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

//Generate UUID
func GenerateCryptoID() string {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        panic(err)
    }
    return hex.EncodeToString(bytes)
}