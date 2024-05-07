package main

import (
	"context"
	"fmt"
	"log"
	"os"
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
