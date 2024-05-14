package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var rootPfad string
var rootPfadOutput string
var rootPfadInput string
var rootPfadDatenbank string
var wg sync.WaitGroup
var pfaden [][]string

func (a *App) BlameStartup() bool {
	start := time.Now()
	rootPfad = "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\"
	rootPfadOutput = rootPfad + "BlameOutput\\"
	rootPfadInput = rootPfad + "BlameInput\\"
	rootPfadDatenbank = rootPfad + "BlameDatenbank\\"
	// Code to measure
	pfaden = [][]string{
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Lagerhueter.xlsx", "KNT", "Lager"},
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Topix.xlsx", "SITECA", "Lager"},
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Moeller.xlsx", "MOELLER", "Lager"},
	}

	for _, pfad := range pfaden {

		pfadLen := len(strings.Split(pfad[0], "\\"))
		fileNameVoll := strings.Split(pfad[0], "\\")[pfadLen-1]
		fileName := strings.Split(fileNameVoll, ".")[0]
		fileExtension := strings.Split(fileNameVoll, ".")[1]
		fmt.Println(fileNameVoll)
		fmt.Println(fileName)
		fmt.Println(fileExtension)

		wg.Add(1)
		ImportFile(pfad[0], pfad[1], pfad[2], fileName)
		wg.Add(1)
		loadFile(pfad[1], pfad[2], fileName)
	}

	//INIT_VERBINDUNGSLITE()
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
	return true
}

func (a *App) LoadStueckliste(pfad []string, kunde string, fileType string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	start := time.Now()
	for _, pfad2 := range pfad {
		pfadLen := len(strings.Split(pfad2, "\\"))
		fileNameVoll := strings.Split(pfad2, "\\")[pfadLen-1]
		fileName := strings.Split(fileNameVoll, ".")[0]
		fileExtension := strings.Split(fileNameVoll, ".")[1]
		fmt.Println(fileNameVoll)
		fmt.Println(fileName)
		fmt.Println(fileExtension)

		wg.Add(1)
		fmt.Println("Importing " + fileName)
		ImportFile(pfad2, kunde, fileType, fileName)
		wg.Add(1)
		fmt.Println("Loading " + fileName)
		loadFile(kunde, fileType, fileName)
		wg.Add(1)
		fmt.Println("Summing " + fileName)
		sumListe(kunde, fileType, fileName)

	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
}

var currentProject *PROJEKT

func (a *App) NewProject2() {

	currentProject = &PROJEKT{
		PROJEKT_NUMMER:       "8000772",
		PROJEKT_BESCHREIBUNG: "Polifilm",
		BAUJAHR:              2024,
		AKTIV:                true,
	}

	b2, err := json.MarshalIndent(currentProject, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(b2))
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\.blame.json", b2, 0644)
	if err != nil {
		log.Println(err)
	}
}
