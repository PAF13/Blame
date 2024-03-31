package main

import (
	"context"
	"fmt"
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
	result, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Question",
	})
	return result
}

func (a *App) ExcelChoice(file1 string, file2 string) {
	Compare(ImportXLSX(file1), ImportXLSX(file2))
	return
}
func ImportXLSX(x string) map[string]int {
	m := make(map[string]int)
	headSkip := 1
	skip := 0
	//open file
	f, err := excelize.OpenFile(x)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//read rows in sheet
	rows, err := f.GetRows("Blatt1")
	if err != nil {
		fmt.Println(err)
	}

	//convert rows into string int
	for _, row := range rows {
		if skip > headSkip {
			//fmt.Print(row[6], "\t")
			//fmt.Println(row[7], "\t")
			rowint, _ := strconv.Atoi(row[7])
			m[row[6]] = rowint
		}
		skip++

	}

	return m
}

func Compare(x map[string]int, y map[string]int) {
	listDif := make(map[string]int)
	for k, v := range y {
		_, ok := x[k]
		if ok {
			//add changed items
			if v != x[k] {
				listDif[k] = v - x[k]
				fmt.Println("Dif |", k, ":", listDif[k])
				delete(x, k)
			}

		} else {
			//add new items
			listDif[k] = v
		}
		delete(y, k)
	}
	for k, v := range x {

		listDif[k] = v * -1
		fmt.Println("Dif |", k, ":", listDif[k])
		delete(x, k)
	}
}
