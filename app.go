package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test(file1 string) string {
	return fmt.Sprintf("Hello %s, It's show time!", file1)
}

func (a *App) Message(file1 string) string {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Question",
		Message:       "Do you want to continue?",
		DefaultButton: "No",
	})
	return result
}

func (a *App) Dialog() string {
	fmt.Println("Dialog start")
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return result
}

func (a *App) ListTrees() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func (a *App) LoadCSV(x string) {
}

func (a *App) NewProject(Jahr string, Kunde string, Projektname string) {
	Projektpfad := "\\\\ME-Datenbank-1\\Projektdaten " + Jahr + "\\" + Kunde + "\\"
	err := os.Mkdir(Projektpfad+Projektname, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Projektpfad + Projektname)
	Projektinfodatei, err := os.Create(Projektpfad + Projektname + "\\" + ".Blame")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Projektinfodatei)
	Projektinfodatei.Close()
}

// stücklistevergleich
func (a *App) ExcelChoice(file1 string, file2 string) {
	fmt.Println("Stueckliste Compare: Recieving Stücklisten")
	CompareStueckliste(LoadStueckliste(file1), LoadStueckliste(file2))
	fmt.Println("Stueckliste Compare: Stücklisten recieved")
}
func LoadStueckliste(x string) map[string][]string {
	fmt.Println("Stueckliste Compare: Creating map for:", x)
	//creating map of list
	stuecklisteMap := make(map[string][]string)
	headSkip := 1
	skip := 0
	//opening spreadsheet
	spreadsheet, err := excelize.OpenFile(x)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := spreadsheet.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(spreadsheet.GetSheetList()[0])
	//Selecting rows from first sheet
	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}
	//reading rows
	for line, row := range rows {
		//cleaning empty erp numbers
		if row[6] == "" {
			t := strconv.Itoa(line)
			row[6] = "Empty" + t
		}
		//loading map with data as string
		if skip > headSkip {
			stuecklisteMap[row[6]] = row
			//fmt.Println(stuecklisteMap[row[6]])
		}
		skip++
	}
	return stuecklisteMap
}
func CompareStueckliste(old map[string][]string, new map[string][]string) {
	file := excelize.NewFile()
	headers := []string{
		"Artikelnummer",
		"»»» Stücklisten/Sets «««",
		"ist Stückliste",
		"Stücklistenart",
		"Positionen ausblenden",
		"SL-Pos.Rang",
		"SL-Pos.Nummer",
		"SL-Pos.Menge",
		"Löschen",
	}
	headers2 := []string{
		"Stücklisten-Kopfartikel, dieser muß schon in T8 angelegt sein.",
		"»»» Stücklisten/Sets «««",
		"1=JA",
		"0=Kopf ohne Positionen, 1=Pos. ohne Preise, 2=Pos mit Preisen",
		"A, B, L ,R",
		"GANZ WICHTIG: durchgehend nummerieren, sonst werden keine neuen Positionen angefügt",
		"Artikelnummer des zugehörigen Artikels",
		"Stücklisten-menge",
		"Hersteller",
		"Typnummer",
		"Artikelnummer",
		"Artikel: Bezeichnung",
	}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for i, header := range headers2 {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 2), header)
	}
	line := 3
	for newValue, _ := range new {
		_, Match := old[newValue]
		if Match && new[newValue][7] != old[newValue][7] {
			for i := 6; i < len(new[newValue]); i++ {
				if i == 7 {
					mengeOld, _ := strconv.Atoi(old[newValue][i])
					mengeNew, _ := strconv.Atoi(new[newValue][i])
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew-mengeOld)
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), new[newValue][i])
				}
			}
			line++
		} else if Match == false {
			for i := 6; i < len(new[newValue]); i++ {
				if i == 7 {
					mengeNew, _ := strconv.Atoi(new[newValue][i])
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew)
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), new[newValue][i])
				}
			}
			delete(new, newValue)
			line++
		}
		delete(old, newValue)

	}
	for newValue, _ := range old {
		for i := 6; i < len(old[newValue]); i++ {
			if i == 7 {
				mengeNew, _ := strconv.Atoi(old[newValue][i])
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew*-1)
			} else {
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), old[newValue][i])
			}

		}
		line++
		delete(old, newValue)
	}

	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Stueckliste.xlsx"); err != nil {
		fmt.Println(err)
	}
}
